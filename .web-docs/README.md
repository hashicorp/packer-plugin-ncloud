The ncloud plugin allows you to create custom images on the NAVER Cloud Platform.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    ncloud = {
      source  = "github.com/hashicorp/ncloud"
      version = "~> 1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/hashicorp/ncloud
```

### Components

#### Builders

- [ncloud](/packer/integrations/hashicorp/ncloud/latest/components/builder/ncloud) - The `ncloud` builder allows you to create server images using the
  [NAVER Cloud Platform](https://www.ncloud.com/).
