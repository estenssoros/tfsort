# tfsort

A command-line tool for sorting Terraform blocks alphabetically by name. Currently supports sorting `variable` and `output` blocks in Terraform files.

## Features

- ✅ Sort Terraform variable blocks alphabetically by name
- ✅ Sort Terraform output blocks alphabetically by name  
- ✅ Preserves original formatting and comments
- ✅ Works with standard `.tf` files
- ✅ Fast HCL parsing using HashiCorp's official parser

## Installation

### Prerequisites

- Go 1.24.2 or later

### Install from source

```bash
git clone https://github.com/estenssoros/tfsort.git
cd tfsort
go install
```

## Usage

```bash
tfsort <command> <file.tf>
```

### Commands

#### `variables`
Sort variable blocks in a Terraform file alphabetically by variable name.

```bash
tfsort variables terraform_file.tf
```

#### `outputs`
Sort output blocks in a Terraform file alphabetically by output name.

```bash
tfsort outputs terraform_file.tf
```

## Examples

### Sorting Variables

Given a Terraform file `variables.tf` with unsorted variables:

```hcl
variable "zebra" {
  description = "A zebra"
  type        = string
}

variable "apple" {
  description = "An apple"
  type        = string
}

variable "banana" {
  description = "A banana"
  type        = string
}
```

Running `tfsort variables variables.tf` will output:

```hcl
variable "apple" {
  description = "An apple"
  type        = string
}

variable "banana" {
  description = "A banana"
  type        = string
}

variable "zebra" {
  description = "A zebra"
  type        = string
}
```

### Sorting Outputs

Given a Terraform file `outputs.tf`:

```hcl
output "vpc_id" {
  description = "ID of the VPC"
  value       = aws_vpc.main.id
}

output "availability_zones" {
  description = "List of availability zones"
  value       = data.aws_availability_zones.available.names
}

output "subnet_ids" {
  description = "List of subnet IDs"
  value       = aws_subnet.main[*].id
}
```

Running `tfsort outputs outputs.tf` will output the blocks sorted alphabetically:

```hcl
output "availability_zones" {
  description = "List of availability zones"
  value       = data.aws_availability_zones.available.names
}

output "subnet_ids" {
  description = "List of subnet IDs"
  value       = aws_subnet.main[*].id
}

output "vpc_id" {
  description = "ID of the VPC"
  value       = aws_vpc.main.id
}
```

## Common Usage Patterns

### Overwrite file with sorted content

```bash
tfsort variables variables.tf > variables_sorted.tf
mv variables_sorted.tf variables.tf
```

### Sort multiple files

```bash
for file in *.tf; do
  tfsort variables "$file" > "${file}.sorted"
  mv "${file}.sorted" "$file"
done
```

### Preview changes before applying

```bash
# See what the sorted output would look like
tfsort variables variables.tf

# Compare with original
diff variables.tf <(tfsort variables variables.tf)
```

## Limitations

- Only processes `.tf` files (Terraform configuration files)
- Currently supports only `variable` and `output` blocks
- Outputs to stdout - you need to redirect to overwrite files
- Requires valid HCL syntax to parse correctly

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [HashiCorp HCL](https://github.com/hashicorp/hcl) - HCL parser
- [Logrus](https://github.com/sirupsen/logrus) - Logging
- [pkg/errors](https://github.com/pkg/errors) - Error handling    