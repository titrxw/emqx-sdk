package entity

type ACL_ACTION string

const (
	ACTION_PUB    ACL_ACTION = "pub"
	ACTION_SUB    ACL_ACTION = "sub"
	ACTION_PUBSUB ACL_ACTION = "pubsub"
)

type ACL_ACCESS string

const (
	ACCESS_ALLOW ACL_ACCESS = "allow"
	ACCESS_DENY  ACL_ACCESS = "deny"
)

type AclEntity struct {
	clientName string
	ipaddr     string
	topic      string
	action     ACL_ACTION
	access     ACL_ACCESS
}

func (aclEntity *AclEntity) SetClientName(clientName string) {
	aclEntity.clientName = clientName
}

func (aclEntity *AclEntity) GetClientName() string {
	return aclEntity.clientName
}

func (aclEntity *AclEntity) SetIpaddr(ipaddr string) {
	aclEntity.ipaddr = ipaddr
}

func (aclEntity *AclEntity) GetIpaddr() string {
	return aclEntity.ipaddr
}

func (aclEntity *AclEntity) SetTopic(topic string) {
	aclEntity.topic = topic
}

func (aclEntity *AclEntity) GetTopic() string {
	return aclEntity.topic
}

func (aclEntity *AclEntity) SetAction(action ACL_ACTION) {
	aclEntity.action = action
}

func (aclEntity *AclEntity) SetActionPub() {
	aclEntity.SetAction(ACTION_PUB)
}

func (aclEntity *AclEntity) SetActionSub() {
	aclEntity.SetAction(ACTION_SUB)
}

func (aclEntity *AclEntity) SetActionPubSub() {
	aclEntity.SetAction(ACTION_PUBSUB)
}

func (aclEntity *AclEntity) GetAction() ACL_ACTION {
	return aclEntity.action
}

func (aclEntity *AclEntity) SetAccess(access ACL_ACCESS) {
	aclEntity.access = access
}

func (aclEntity *AclEntity) SetAccessAllow() {
	aclEntity.SetAccess(ACCESS_ALLOW)
}

func (aclEntity *AclEntity) SetAccessDeny() {
	aclEntity.SetAccess(ACCESS_DENY)
}

func (aclEntity *AclEntity) GetAccess() ACL_ACCESS {
	return aclEntity.access
}
