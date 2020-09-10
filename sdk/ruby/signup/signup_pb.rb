# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: signup/signup.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "go.micro.service.signup.SendVerificationEmailRequest" do
    optional :email, :string, 1
  end
  add_message "go.micro.service.signup.SendVerificationEmailResponse" do
  end
  add_message "go.micro.service.signup.VerifyRequest" do
    optional :email, :string, 1
    optional :token, :string, 2
  end
  add_message "go.micro.service.signup.VerifyResponse" do
    optional :authToken, :message, 1, "go.micro.service.signup.AuthToken"
    optional :customerID, :string, 2
    optional :namespace, :string, 3
    optional :message, :string, 4
    optional :payment_required, :bool, 5
    repeated :namespaces, :string, 6
  end
  add_message "go.micro.service.signup.CompleteSignupRequest" do
    optional :email, :string, 1
    optional :token, :string, 2
    optional :paymentMethodID, :string, 3
    optional :secret, :string, 4
    optional :namespace, :string, 5
  end
  add_message "go.micro.service.signup.CompleteSignupResponse" do
    optional :authToken, :message, 1, "go.micro.service.signup.AuthToken"
    optional :namespace, :string, 2
  end
  add_message "go.micro.service.signup.AuthToken" do
    optional :access_token, :string, 1
    optional :refresh_token, :string, 2
    optional :created, :int64, 3
    optional :expiry, :int64, 4
  end
  add_message "go.micro.service.signup.RecoverRequest" do
    optional :email, :string, 1
  end
  add_message "go.micro.service.signup.RecoverResponse" do
  end
end

module Go
  module Micro
    module Service
      module Signup
        SendVerificationEmailRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.SendVerificationEmailRequest").msgclass
        SendVerificationEmailResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.SendVerificationEmailResponse").msgclass
        VerifyRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.VerifyRequest").msgclass
        VerifyResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.VerifyResponse").msgclass
        CompleteSignupRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.CompleteSignupRequest").msgclass
        CompleteSignupResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.CompleteSignupResponse").msgclass
        AuthToken = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.AuthToken").msgclass
        RecoverRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.RecoverRequest").msgclass
        RecoverResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("go.micro.service.signup.RecoverResponse").msgclass
      end
    end
  end
end
