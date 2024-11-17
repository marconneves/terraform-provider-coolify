# Provider Deprecation Notice

This repository is being discontinued in favor of the [terraform-provider-coolify](https://github.com/SierraJC/terraform-provider-coolify), which is actively maintained and supports the same functionality. We recommend using the new provider to ensure compatibility and continued support.

## Why the Change?

The decision to deprecate this provider was made to consolidate efforts and offer a more robust solution. The `terraform-provider-coolify` provides comprehensive support for managing resources in Coolify, and it is being continuously updated by the maintainers.

## What You Need to Do

If you're currently using this provider, we encourage you to migrate to `terraform-provider-coolify` as soon as possible. Here's how you can transition:

1. **Review the Documentation**:  
   Familiarize yourself with the [terraform-provider-coolify documentation](https://github.com/SierraJC/terraform-provider-coolify) to understand the available features and configurations.

2. **Check Compatibility**:  
   Ensure that `terraform-provider-coolify` is compatible with your current Terraform version. The provider's documentation will specify the supported versions.

3. **Update Your Configuration**:  
   Replace this provider in your Terraform configuration files with `terraform-provider-coolify`. You will need to update the `provider` block and modify resource configurations as needed.

4. **Test in a Development Environment**:  
   Before applying changes to production, test your configuration in a development environment to verify that everything works as expected.

5. **Consult the Community**:  
   If you have questions or encounter issues, visit the [Issues section](https://github.com/SierraJC/terraform-provider-coolify/issues) of the `terraform-provider-coolify` repository or consult relevant forums and communities.

## Next Steps

For any new development or issues, please direct them to the [terraform-provider-coolify repository](https://github.com/SierraJC/terraform-provider-coolify). 

Thank you for your understanding, and we hope the transition is smooth.