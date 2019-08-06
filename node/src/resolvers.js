module.exports = {
  Query: {
    viewer,
    contracts,
  },
  Viewer: {
    notifications,
  },
  Notification: {
    kind: (parent) => parent.kind.toUpperCase()
  }
};

function viewer() {
  return {
    id: "1",
    name: 'Samuel Node',
    avatar: {
      title: 'Samuel Avatar',
      href: 'https://avatar.com/1.jpg',
    }
  }
}

const ContractStatus = { ACTIVE: 'active', INACTIVE: 'inactive' };

const _contracts = [
  {originalId: "1", amount: 11.11, status: ContractStatus.ACTIVE, ownerId: "1" },
  {originalId: "2", amount: 22.22, status: ContractStatus.ACTIVE, ownerId: "2" },
  {originalId: "3", amount: 33.33, status: ContractStatus.INACTIVE, ownerId: "1" },
  {originalId: "4", amount: 44.44, status: ContractStatus.INACTIVE, ownerId: "2" },
]

function contracts(_, args) {
  if (args.filter && args.filter.outer === 'ALL') {
    return {
      totalCount: _contracts.length,
      nodes: _contracts,
    };
  }

  const nodes = _contracts.filter(c => c.ownerId === "1");
  return {
    nodes,
    totalCount: nodes.length,
  }
}

const NotificationKind = {
  INFO: 'info',
  ERROR: 'error',
  WARNING: 'warning',
}

const _notifications = [
  { ownerId: "2", title: "i am a error", kind: NotificationKind.ERROR },
  { ownerId: "2", title: "i am a warning", kind: NotificationKind.WARNING },
  { ownerId: "1", title: "i am a info", kind: NotificationKind.INFO },
]

function notifications(parent) {
  return _notifications.filter(n => n.ownerId === parent.id);
}
