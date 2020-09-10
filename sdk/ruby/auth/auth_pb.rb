# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: auth/auth.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "auth.ListAccountsRequest" do
    optional :options, :message, 1, "auth.Options"
  end
  add_message "auth.ListAccountsResponse" do
    repeated :accounts, :message, 1, "auth.Account"
  end
  add_message "auth.DeleteAccountRequest" do
    optional :id, :string, 1
    optional :options, :message, 2, "auth.Options"
  end
  add_message "auth.DeleteAccountResponse" do
  end
  add_message "auth.Token" do
    optional :access_token, :string, 1
    optional :refresh_token, :string, 2
    optional :created, :int64, 3
    optional :expiry, :int64, 4
  end
  add_message "auth.Account" do
    optional :id, :string, 1
    optional :type, :string, 2
    map :metadata, :string, :string, 4
    repeated :scopes, :string, 5
    optional :issuer, :string, 6
    optional :secret, :string, 7
  end
  add_message "auth.Resource" do
    optional :name, :string, 1
    optional :type, :string, 2
    optional :endpoint, :string, 3
  end
  add_message "auth.GenerateRequest" do
    optional :id, :string, 1
    map :metadata, :string, :string, 3
    repeated :scopes, :string, 4
    optional :secret, :string, 5
    optional :type, :string, 6
    optional :provider, :string, 7
    optional :options, :message, 8, "auth.Options"
  end
  add_message "auth.GenerateResponse" do
    optional :account, :message, 1, "auth.Account"
  end
  add_message "auth.GrantRequest" do
    optional :scope, :string, 1
    optional :resource, :message, 2, "auth.Resource"
    optional :options, :message, 3, "auth.Options"
  end
  add_message "auth.GrantResponse" do
  end
  add_message "auth.RevokeRequest" do
    optional :scope, :string, 1
    optional :resource, :message, 2, "auth.Resource"
    optional :options, :message, 3, "auth.Options"
  end
  add_message "auth.RevokeResponse" do
  end
  add_message "auth.InspectRequest" do
    optional :token, :string, 1
    optional :options, :message, 2, "auth.Options"
  end
  add_message "auth.InspectResponse" do
    optional :account, :message, 1, "auth.Account"
  end
  add_message "auth.TokenRequest" do
    optional :id, :string, 1
    optional :secret, :string, 2
    optional :refresh_token, :string, 3
    optional :token_expiry, :int64, 4
    optional :options, :message, 5, "auth.Options"
  end
  add_message "auth.TokenResponse" do
    optional :token, :message, 1, "auth.Token"
  end
  add_message "auth.Rule" do
    optional :id, :string, 1
    optional :scope, :string, 2
    optional :resource, :message, 3, "auth.Resource"
    optional :access, :enum, 4, "auth.Access"
    optional :priority, :int32, 5
  end
  add_message "auth.Options" do
    optional :namespace, :string, 1
  end
  add_message "auth.CreateRequest" do
    optional :rule, :message, 1, "auth.Rule"
    optional :options, :message, 2, "auth.Options"
  end
  add_message "auth.CreateResponse" do
  end
  add_message "auth.DeleteRequest" do
    optional :id, :string, 1
    optional :options, :message, 2, "auth.Options"
  end
  add_message "auth.DeleteResponse" do
  end
  add_message "auth.ListRequest" do
    optional :options, :message, 2, "auth.Options"
  end
  add_message "auth.ListResponse" do
    repeated :rules, :message, 1, "auth.Rule"
  end
  add_message "auth.ChangeSecretRequest" do
    optional :id, :string, 1
    optional :old_secret, :string, 2
    optional :new_secret, :string, 3
    optional :options, :message, 4, "auth.Options"
  end
  add_message "auth.ChangeSecretResponse" do
  end
  add_enum "auth.Access" do
    value :UNKNOWN, 0
    value :GRANTED, 1
    value :DENIED, 2
  end
end

module Auth
  ListAccountsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.ListAccountsRequest").msgclass
  ListAccountsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.ListAccountsResponse").msgclass
  DeleteAccountRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.DeleteAccountRequest").msgclass
  DeleteAccountResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.DeleteAccountResponse").msgclass
  Token = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.Token").msgclass
  Account = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.Account").msgclass
  Resource = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.Resource").msgclass
  GenerateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.GenerateRequest").msgclass
  GenerateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.GenerateResponse").msgclass
  GrantRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.GrantRequest").msgclass
  GrantResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.GrantResponse").msgclass
  RevokeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.RevokeRequest").msgclass
  RevokeResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.RevokeResponse").msgclass
  InspectRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.InspectRequest").msgclass
  InspectResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.InspectResponse").msgclass
  TokenRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.TokenRequest").msgclass
  TokenResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.TokenResponse").msgclass
  Rule = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.Rule").msgclass
  Options = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.Options").msgclass
  CreateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.CreateRequest").msgclass
  CreateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.CreateResponse").msgclass
  DeleteRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.DeleteRequest").msgclass
  DeleteResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.DeleteResponse").msgclass
  ListRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.ListRequest").msgclass
  ListResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.ListResponse").msgclass
  ChangeSecretRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.ChangeSecretRequest").msgclass
  ChangeSecretResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.ChangeSecretResponse").msgclass
  Access = Google::Protobuf::DescriptorPool.generated_pool.lookup("auth.Access").enummodule
end
