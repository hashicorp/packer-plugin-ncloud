# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Naver Cloud"
  description = "The Naver Cloud plugin can be used with HashiCorp Packer to create custom images on the Naver Cloud platform"
  identifier = "packer/hashicorp/ncloud"
  component {
    type = "builder"
    name = "Naver Cloud Platform"
    slug = "ncloud"
  }
}
