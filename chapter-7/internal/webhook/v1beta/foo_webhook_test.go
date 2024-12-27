/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	samplecontrollerv1beta "github.com/naruse666/road-to-custom-controller/chapter-7/api/v1beta"
	// TODO (user): Add any additional imports if needed
)

var _ = Describe("Foo Webhook", func() {
	var (
		obj    *samplecontrollerv1beta.Foo
		oldObj *samplecontrollerv1beta.Foo
	)

	BeforeEach(func() {
		obj = &samplecontrollerv1beta.Foo{}
		oldObj = &samplecontrollerv1beta.Foo{}
		Expect(oldObj).NotTo(BeNil(), "Expected oldObj to be initialized")
		Expect(obj).NotTo(BeNil(), "Expected obj to be initialized")
		// TODO (user): Add any setup logic common to all tests
	})

	AfterEach(func() {
		// TODO (user): Add any teardown logic common to all tests
	})

	Context("When creating Foo under Conversion Webhook", func() {
		// TODO (user): Add logic to convert the object to the desired version and verify the conversion
		// Example:
		// It("Should convert the object correctly", func() {
		//     convertedObj := &samplecontrollerv1beta.Foo{}
		//     Expect(obj.ConvertTo(convertedObj)).To(Succeed())
		//     Expect(convertedObj).ToNot(BeNil())
		// })
	})

})
