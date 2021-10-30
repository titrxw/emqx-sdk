package acl

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
	userName string
	clientId string
	ipaddr   string
	topic    string
	action   ACL_ACTION
	access   ACL_ACCESS
}

func (aclEntity *AclEntity) setUserName(userName string) {
	aclEntity.userName = userName
}

func (aclEntity *AclEntity) getUserName() string {
	return aclEntity.userName
}

func (aclEntity *AclEntity) setClientId(clientId string) {
	aclEntity.clientId = clientId
}

func (aclEntity *AclEntity) getClientId() string {
	return aclEntity.clientId
}

func (aclEntity *AclEntity) setIpaddr(ipaddr string) {
	aclEntity.ipaddr = ipaddr
}

func (aclEntity *AclEntity) getIpaddr() string {
	return aclEntity.ipaddr
}

func (aclEntity *AclEntity) setTopic(topic string) {
	aclEntity.topic = topic
}

func (aclEntity *AclEntity) getTopic() string {
	return aclEntity.topic
}

func (aclEntity *AclEntity) setActionPub() {
	aclEntity.action = ACTION_PUB
}

func (aclEntity *AclEntity) setActionSub() {
	aclEntity.action = ACTION_SUB
}

func (aclEntity *AclEntity) setActionPubSub() {
	aclEntity.action = ACTION_PUBSUB
}

func (aclEntity *AclEntity) getAction() ACL_ACTION {
	return aclEntity.action
}

func (aclEntity *AclEntity) setAccessAllow() {
	aclEntity.access = ACCESS_ALLOW
}

func (aclEntity *AclEntity) setAccessDeny() {
	aclEntity.access = ACCESS_DENY
}

func (aclEntity *AclEntity) getAccess() ACL_ACCESS {
	return aclEntity.access
}
