
locals {
  cluster_name = "littlelink"
}

# EKS Cluster IAM Role
resource "aws_iam_role" "eks_cluster_role" {
  name = "${local.cluster_name}-cluster-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "eks.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}


# Attach the AmazonEKSClusterPolicy to the role
resource "aws_iam_role_policy_attachment" "eks_cluster_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.eks_cluster_role.name
}

# EKS Node IAM Role
resource "aws_iam_role" "eks_node_role" {
  name = "${local.cluster_name}-node-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ec2.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

# Attach necessary policies for the worker nodes
resource "aws_iam_role_policy_attachment" "eks_worker_node_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
  role       = aws_iam_role.eks_node_role.name
}

resource "aws_iam_role_policy_attachment" "eks_cni_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
  role       = aws_iam_role.eks_node_role.name
}

resource "aws_iam_role_policy_attachment" "ec2_container_registry_read_only" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
  role       = aws_iam_role.eks_node_role.name
}


# --- EKS Cluster ---

# Retrieve available availability zones
data "aws_availability_zones" "available" {}

# Use the official Terraform EKS module
module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "~> 20.0" # Use a specific version for production

  cluster_name    = local.cluster_name
  cluster_version = "1.33"

  cluster_endpoint_public_access = true


  vpc_id     = aws_default_vpc.default.id
  subnet_ids = [aws_default_subnet.public_subnet1.id, aws_default_subnet.public_subnet2.id, aws_default_subnet.public_subnet3.id]

  # --- Managed Node Group for EC2 instances ---
  eks_managed_node_groups = {
    main = {
      name           = "${local.cluster_name}-node-group"
      instance_types = ["t3.medium"]
      min_size       = 1
      max_size       = 3
      desired_size   = 2

      iam_role_arn = aws_iam_role.eks_node_role.arn
    }
  }

  tags = {
    Environment = "development"
    Project     = "my-eks-project"
  }
  access_entries = {
    my_admin_role = {
      kubernetes_groups = []
      principal_arn     = "arn:aws:iam::247487031771:user/adminuser" # Or user ARN

      policy_associations = {
        my_admin_role = {
          policy_arn    = "arn:aws:eks::aws:cluster-access-policy/AmazonEKSClusterAdminPolicy"
          principal_arn = "arn:aws:iam::247487031771:user/adminuser"
          access_scope = {
            type = "cluster"
          }
        }
      }
    }
  }
}

# --- Outputs ---

output "cluster_endpoint" {
  description = "Endpoint for EKS cluster."
  value       = module.eks.cluster_endpoint
}

output "cluster_name" {
  description = "Name of the EKS cluster."
  value       = module.eks.cluster_name
}

# output "kubeconfig_command" {
#   description = "Command to configure kubectl."
#   value       = "aws eks update-kubeconfig --region ${var.region} --name ${module.eks.cluster_name}"
# }
