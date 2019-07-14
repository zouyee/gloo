package awscache

import (
	"context"
	"fmt"
	"strings"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws/glooec2/utils"

	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/ec2/awslister"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws/glooec2"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Batcher tests", func() {
	It("synopsis, basic check", func() {
		secretMeta1 := core.Metadata{Name: "s1", Namespace: "default"}
		secretRef1 := secretMeta1.Ref()
		secret1 := &v1.Secret{
			Kind: &v1.Secret_Aws{
				Aws: &v1.AwsSecret{
					AccessKey: "very-access",
					SecretKey: "very-secret",
				},
			},
			Metadata: secretMeta1,
		}
		secrets := v1.SecretList{secret1}
		cb := newCache(context.TODO(), secrets)
		region1 := "us-east-1"
		upstream := &v1.Upstream{
			UpstreamSpec: &v1.UpstreamSpec{
				UpstreamType: &v1.UpstreamSpec_AwsEc2{
					AwsEc2: &glooec2.UpstreamSpec{
						Region:    region1,
						SecretRef: secretRef1,
						Filters: []*glooec2.TagFilter{{
							Spec: &glooec2.TagFilter_Key{
								Key: "k1",
							},
						}},
						PublicIp: false,
						Port:     8080,
					},
				},
			},
			Status:   core.Status{},
			Metadata: core.Metadata{},
		}
		up1 := invertKnownEc2Upstream(upstream)
		err := cb.addUpstream(up1)
		Expect(err).NotTo(HaveOccurred())

		credSpec1 := awslister.NewCredentialSpec(secretRef1, region1, nil)
		instances := []*ec2.Instance{{
			Tags: []*ec2.Tag{{
				Key:   aws.String("k1"),
				Value: aws.String("any old value"),
			}},
		}}
		Expect(cb.addInstances(credSpec1, instances)).NotTo(HaveOccurred())
		filteredInstances1, err := cb.FilterEndpointsForUpstream(up1)
		Expect(err).NotTo(HaveOccurred())
		Expect(filteredInstances1).To(Equal(instances))

	})

	// Represent 3 credential specification cases:
	// A: secret has full access to credentialMap in it region
	// B: secret has limited access to credentialMap in it region
	// C: same as A, different region
	var (
		region1             = "us-east-1"
		region2             = "us-east-2"
		instance1           = generateTestInstance("1") // region1
		instance2           = generateTestInstance("2") // region1
		instance3           = generateTestInstance("3") // region1
		instance4           = generateTestInstance("4") // region2
		instance5           = generateTestInstance("5") // region2
		secret1, secretRef1 = generateTestSecrets("1")
		secret2, secretRef2 = generateTestSecrets("2")
		credSpecA           = generateCredSpec(region1, secretRef1)
		credSpecB           = generateCredSpec(region1, secretRef2)
		credSpecC           = generateCredSpec(region2, secretRef2)
		credInstancesA      = []*ec2.Instance{instance1, instance2, instance3}
		credInstancesB      = []*ec2.Instance{instance1, instance2}
		credInstancesC      = []*ec2.Instance{instance4, instance5}
	)
	// In this table test, we use a fixed set of instances and secrets
	// Each table entry describes what filters should be applied to the upstream and what instances should be returned
	DescribeTable("cache should assemble and disassemble credential-grouped results", func(input filterTestInput) {
		secrets := v1.SecretList{secret1, secret2}
		cb := newCache(context.TODO(), secrets)

		// build the dummy upstreams
		upA := generateUpstreamWithCredentials("A", credSpecA, filterTestInput{})
		upB := generateUpstreamWithCredentials("B", credSpecB, filterTestInput{})
		upC := generateUpstreamWithCredentials("C", credSpecC, filterTestInput{})
		// build the upstream that we care about
		upTest := generateUpstreamWithCredentials("Test", input.credentialSpec, input)
		// prime the map with the upstreams
		Expect(cb.addUpstream(upA)).NotTo(HaveOccurred())
		Expect(cb.addUpstream(upB)).NotTo(HaveOccurred())
		Expect(cb.addUpstream(upC)).NotTo(HaveOccurred())
		Expect(cb.addUpstream(upTest)).NotTo(HaveOccurred())

		// "query" the api for each upstream
		Expect(cb.addInstances(credSpecA, credInstancesA)).NotTo(HaveOccurred())
		Expect(cb.addInstances(credSpecB, credInstancesB)).NotTo(HaveOccurred())
		Expect(cb.addInstances(credSpecC, credInstancesC)).NotTo(HaveOccurred())

		// core test: apply the filter, assert expectations
		filteredInstances, err := cb.FilterEndpointsForUpstream(upTest)
		Expect(err).NotTo(HaveOccurred())
		//Expect(len(filteredInstances)).To(Equal(len(input.expected)))
		var filteredIds []string
		for _, instance := range filteredInstances {
			filteredIds = append(filteredIds, aws.StringValue(instance.InstanceId))
		}
		var expectedIds []string
		for _, instance := range input.expected {
			expectedIds = append(expectedIds, aws.StringValue(instance.InstanceId))
		}
		Expect(filteredIds).To(ConsistOf(expectedIds))
		Expect(filteredInstances).To(ConsistOf(input.expected))

	},
		Entry("get every instance in region 1", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      nil,
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance1, instance2, instance3},
		}),
		Entry("get instance 1 in region 1", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      []string{"k1a"},
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance1},
		}),
		Entry("not match with unused tag key", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      []string{tagKeyThatIsNotUsed},
			keyValueFilters: nil,
			expected:        nil,
		}),
		Entry("get multiple instances having common tags", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      []string{commonKeyKey},
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance1, instance2, instance3},
		}),
		Entry("get multiple instances having common tags", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      []string{commonTagKeyKey},
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance1, instance2, instance3},
		}),
		Entry("match with key and value", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      nil,
			keyValueFilters: []string{"k1a:v1a"},
			expected:        []*ec2.Instance{instance1},
		}),
		Entry("match with multiple keys and values", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      nil,
			keyValueFilters: []string{"k1a:v1a", "k1b:v1b"},
			expected:        []*ec2.Instance{instance1},
		}),
		Entry("not match any when filters are too restrictive", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      nil,
			keyValueFilters: []string{"k1a:v1a", "akey:not_in_use"},
			expected:        nil,
		}),
		// casing
		Entry("key case does not matter for key filters", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      []string{"K1A", "k1B"},
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance1},
		}),
		Entry("key case does not matter for key value filter keys", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      nil,
			keyValueFilters: []string{"K1A:v1a", "K1b:v1b"},
			expected:        []*ec2.Instance{instance1},
		}),
		Entry("key case does matter for key value filter values", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      nil,
			keyValueFilters: []string{"k1a:V1A"},
			expected:        nil,
		}),
		// misc. test consistency validation
		Entry("test works with other credentials, no filters", filterTestInput{
			credentialSpec:  credSpecB,
			keyFilters:      nil,
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance1, instance2},
		}),
		Entry("test works with other credentials, k filters", filterTestInput{
			credentialSpec:  credSpecB,
			keyFilters:      []string{"k1a"},
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance1},
		}),
		Entry("test works with other credentials, kv filters", filterTestInput{
			credentialSpec:  credSpecB,
			keyFilters:      nil,
			keyValueFilters: []string{"k2a:v2a"},
			expected:        []*ec2.Instance{instance2},
		}),
		Entry("test works with other region, no filters", filterTestInput{
			credentialSpec:  credSpecC,
			keyFilters:      nil,
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance4, instance5},
		}),
		Entry("test works with other region, k filters", filterTestInput{
			credentialSpec:  credSpecC,
			keyFilters:      []string{"k4b"},
			keyValueFilters: nil,
			expected:        []*ec2.Instance{instance4},
		}),
		Entry("test works with other region, kv filters", filterTestInput{
			credentialSpec:  credSpecC,
			keyFilters:      nil,
			keyValueFilters: []string{"k5b:v5b"},
			expected:        []*ec2.Instance{instance5},
		}),
		Entry("credA returns no matches for (irrelevant) filters from credC", filterTestInput{
			credentialSpec:  credSpecA,
			keyFilters:      nil,
			keyValueFilters: []string{"k5b:v5b"},
			expected:        nil,
		}))
})

var _ = Describe("constructor tests", func() {
	It("basic construction from New", func() {
		upstream := getAnUpstream()
		instance := getAnInstance()
		upstreams := utils.BuildInvertedUpstreamRefMap(v1.UpstreamList{upstream})
		iUpstream := upstreams[upstream.Metadata.Ref()]

		responses := getMockListerResponses(iUpstream)
		mockLister := newMockEc2InstanceLister(responses)
		resp := make(mockListerResponses)
		cspec1 := awslister.NewCredentialSpecFromEc2UpstreamSpec(iUpstream.AwsEc2Spec)
		instances := []*ec2.Instance{instance}
		resp[cspec1.GetKey()] = instances
		secretMeta1 := core.Metadata{Name: "secret1", Namespace: "namespace"}
		secret1 := &v1.Secret{
			Kind: &v1.Secret_Aws{
				Aws: &v1.AwsSecret{
					AccessKey: "very-access",
					SecretKey: "very-secret",
				},
			},
			Metadata: secretMeta1,
		}
		secrets := v1.SecretList{secret1}
		cb, err := New(context.TODO(), secrets, upstreams, mockLister)
		Expect(err).NotTo(HaveOccurred())
		filteredInstances1, err := cb.FilterEndpointsForUpstream(iUpstream)
		Expect(err).NotTo(HaveOccurred())
		Expect(filteredInstances1).To(Equal(instances))

	})
})

type filterTestInput struct {
	// use these credentials when accessing
	credentialSpec *awslister.CredentialSpec
	// format: <key>
	keyFilters []string
	// format: <key>:<value>
	keyValueFilters []string
	// these instances should be returned
	expected []*ec2.Instance
}

const (
	commonKeyKey        = "common_key_only"
	commonTagKeyKey     = "common_key_and_value"
	commonTagKeyValue   = "common_value_d"
	tagKeyThatIsNotUsed = "no_instances_have_this_tag"
)

// outputs basic templated instances for testing various filters
func generateTestInstance(seed string) *ec2.Instance {
	return &ec2.Instance{
		InstanceId: aws.String("i" + seed),
		Tags: []*ec2.Tag{{
			Key:   aws.String("k" + seed + "a"),
			Value: aws.String("v" + seed + "a"),
		}, {
			Key:   aws.String("k" + seed + "b"),
			Value: aws.String("v" + seed + "b"),
		}, {
			Key:   aws.String(commonKeyKey),
			Value: aws.String("v" + seed + "c"),
		}, {
			Key:   aws.String(commonTagKeyKey),
			Value: aws.String(commonTagKeyValue),
		}, {
			Key:   aws.String("unrelated_key_should_not_match"),
			Value: aws.String("unrelated_value_should_not_match"),
		}},
	}

}

func generateTestSecrets(seed string) (*v1.Secret, core.ResourceRef) {
	secretMeta := core.Metadata{Name: "s" + seed, Namespace: "default"}
	secretRef := secretMeta.Ref()
	secret := &v1.Secret{
		Kind: &v1.Secret_Aws{
			Aws: &v1.AwsSecret{
				AccessKey: "abc",
				SecretKey: "123",
			},
		},
		Metadata: secretMeta,
	}
	return secret, secretRef
}

func generateCredSpec(region string, secretRef core.ResourceRef) *awslister.CredentialSpec {
	return awslister.NewCredentialSpec(secretRef, region, nil)

}

// creates an upstream with the filters and credentials defined by the input
func generateUpstreamWithCredentials(name string, credSpec *awslister.CredentialSpec, input filterTestInput) *utils.InvertedEc2Upstream {
	upstreamSpec := &glooec2.UpstreamSpec{
		Region:    credSpec.Region(),
		SecretRef: credSpec.SecretRef(),
	}
	for _, key := range input.keyFilters {
		f := &glooec2.TagFilter{
			Spec: &glooec2.TagFilter_Key{
				Key: key,
			},
		}
		upstreamSpec.Filters = append(upstreamSpec.Filters, f)
	}
	for _, kv := range input.keyValueFilters {
		parts := strings.Split(kv, ":")
		Expect(len(parts)).To(Equal(2))
		key := parts[0]
		val := parts[1]
		f := &glooec2.TagFilter{
			Spec: &glooec2.TagFilter_KvPair_{
				KvPair: &glooec2.TagFilter_KvPair{
					Key:   key,
					Value: val,
				},
			},
		}
		upstreamSpec.Filters = append(upstreamSpec.Filters, f)
	}
	upstream := &v1.Upstream{
		UpstreamSpec: &v1.UpstreamSpec{
			UpstreamType: &v1.UpstreamSpec_AwsEc2{
				AwsEc2: upstreamSpec,
			},
		},
		Status: core.Status{},
		Metadata: core.Metadata{
			Name:      name,
			Namespace: "default",
		},
	}
	inverted := invertKnownEc2Upstream(upstream)
	return inverted
}

type mockListerResponses map[awslister.CredentialKey][]*ec2.Instance
type mockEc2InstanceLister struct {
	responses mockListerResponses
}

func newMockEc2InstanceLister(responses mockListerResponses) *mockEc2InstanceLister {
	// add any test inputs to this
	return &mockEc2InstanceLister{
		responses: responses,
	}
}

func (m *mockEc2InstanceLister) ListForCredentials(ctx context.Context, cred *awslister.CredentialSpec, secrets v1.SecretList) ([]*ec2.Instance, error) {
	v, ok := m.responses[cred.GetKey()]
	if !ok {
		return nil, fmt.Errorf("invalid input, no test responses available")
	}
	return v, nil
}

func getMockListerResponses(iUpstream *utils.InvertedEc2Upstream) mockListerResponses {
	resp := make(mockListerResponses)
	cspec1 := awslister.NewCredentialSpec(iUpstream.AwsEc2Spec.SecretRef, iUpstream.AwsEc2Spec.Region, nil)
	resp[cspec1.GetKey()] = []*ec2.Instance{
		getAnInstance(),
	}
	return resp
}

func getAnInstance() *ec2.Instance {
	var (
		testPrivateIp1 = "111-111-111-111"
		testPublicIp1  = "222.222.222.222"
	)

	return &ec2.Instance{
		PrivateIpAddress: aws.String(testPrivateIp1),
		PublicIpAddress:  aws.String(testPublicIp1),
		Tags: []*ec2.Tag{{
			Key:   aws.String("k1"),
			Value: aws.String("any old value"),
		}},
		VpcId: aws.String("id1"),
	}
}
func getAnUpstream() *v1.Upstream {
	var testPort1 uint32 = 8080
	secretRef := core.ResourceRef{
		Name:      "secret1",
		Namespace: "namespace",
	}
	return &v1.Upstream{
		UpstreamSpec: &v1.UpstreamSpec{
			UpstreamType: &v1.UpstreamSpec_AwsEc2{
				AwsEc2: &glooec2.UpstreamSpec{
					Region:    "us-east-1",
					SecretRef: secretRef,
					Filters: []*glooec2.TagFilter{{
						Spec: &glooec2.TagFilter_Key{
							Key: "k1",
						},
					}},
					PublicIp: false,
					Port:     testPort1,
				},
			}},
		Metadata: core.Metadata{
			Name:      "u1",
			Namespace: "default",
		},
	}
}

func invertKnownEc2Upstream(upstream *v1.Upstream) *utils.InvertedEc2Upstream {
	iMap := utils.BuildInvertedUpstreamRefMap(v1.UpstreamList{upstream})
	inverted, ok := iMap[upstream.Metadata.Ref()]
	Expect(ok).To(BeTrue())
	return inverted
}