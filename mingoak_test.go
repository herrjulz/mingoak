package mingoak_test

import (
	. "github.com/JulzDiverse/mingoak"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mingoak", func() {
	var root *Dir

	BeforeEach(func() {
		root = MkRoot()
	})

	Context("Add/Get Dir", func() {
		Context("When adding an component to a component with Add", func() {
			It("is an existing subcomponent", func() {
				root.MkDirAll("subdir")

				_, err := root.ReadDir("subdir")
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("When adding a component with a path", func() {
			It("it creates all subcomponents", func() {
				root.MkDirAll("path/to/dir/")

				fileInfo, err := root.ReadDir("path")
				Expect(err).ToNot(HaveOccurred())
				Expect(fileInfo[0].Name).To(Equal("to"))
				Expect(len(fileInfo)).To(Equal(1))

				fileInfo, err = root.ReadDir("path/to")
				Expect(fileInfo[0].Name).To(Equal("dir"))
				Expect(err).ToNot(HaveOccurred())

				fileInfo, err = root.ReadDir("path/to/dir")
				Expect(err).ToNot(HaveOccurred())
				Expect(len(fileInfo)).To(Equal(0))

				_, err = root.ReadDir("")
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("When reading an non existing path", func() {
			It("returns an error", func() {
				_, err := root.ReadDir("non_existing")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("Add/Get File", func() {
		Context("When adding a file with Add", func() {
			Context("providing a single file name", func() {

				It("exists in the root", func() {
					root.WriteFile("file", []byte("test"))

					file, err := root.ReadFile("file")
					Expect(err).ToNot(HaveOccurred())

					Expect(string(file)).To(Equal("test"))
				})
			})

			Context("providing a path", func() {
				It("exists in the provided path", func() {
					root.MkDirAll("new/dir/")
					root.WriteFile("new/dir/file", []byte("test"))

					file, err := root.ReadFile("new/dir/file")
					Expect(err).ToNot(HaveOccurred())

					Expect(string(file)).To(Equal("test"))
				})
			})

			Context("providing an empty string", func() {
				It("returns an error", func() {
					root.MkDirAll("new/dir/")
					root.WriteFile("", []byte("test"))

					_, err := root.ReadFile("new/dir/file")
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("File not found!"))
				})
			})
		})
	})
})
