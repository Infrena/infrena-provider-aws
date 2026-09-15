# aws.elasticloadbalancingv2.listener

**CloudFormation type:** `AWS::ElasticLoadBalancingV2::Listener`

Specifies a listener for an Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.

Region attribute: `region`

**Import ID:** `<region>/ListenerArn` (AWS::ElasticLoadBalancingV2::Listener)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AlpnPolicy` | alpn_policy | `list` | optional, computed, provider-chosen |  | [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy. |
| `Certificates` |  | `list` | optional, computed, provider-chosen |  | The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS. |
| `DefaultActions` | default_actions | `list` | required |  | The actions for the default rule. You cannot define a condition for a default rule. |
| `ListenerArn` | listener_arn | `string` | computed |  |  |
| `ListenerAttributes` | listener_attributes | `list` | optional, computed, provider-chosen |  | The listener attributes. Attributes that you do not modify retain their current values. |
| `LoadBalancerArn` | load_balancer_arn | `string` | required, replaces on change | aws.elasticloadbalancingv2.loadbalancer.LoadBalancerArn | The Amazon Resource Name (ARN) of the load balancer. |
| `MutualAuthentication` | mutual_authentication | `map` | optional, computed, provider-chosen |  | The mutual authentication configuration information. |
| `Port` |  | `integer` | optional, computed, provider-chosen |  | The port on which the load balancer is listening. You can't specify a port for a Gateway Load Balancer. |
| `Protocol` |  | `string` | optional, computed, provider-chosen |  | The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, TCP_UDP, QUIC, and TCP_QUIC. You can’t specify the UDP, TCP_UDP, QUIC, or TCP_QUIC protocol if dual-stack mode is enabled. You can't specify a protocol for a Gateway Load Balancer. |
| `SslPolicy` | ssl_policy | `string` | optional, computed, provider-chosen |  | [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported. For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/describe-ssl-policies.html) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/describe-ssl-policies.html) in the *Network Load Balancers Guide*. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
