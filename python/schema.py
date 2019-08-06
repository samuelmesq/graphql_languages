import graphene

OrderByDirection = graphene.Enum('OrderByDirection', ['ASC', 'DESC'])

class OrderBy(graphene.InputObjectType):
    field = graphene.String(required=True)
    direction = graphene.Field(OrderByDirection, required=True)

class Image(graphene.ObjectType):
    href = graphene.String(required=True)
    title = graphene.String(required=True)

class NotificationKind(graphene.Enum):
    INFO = 0
    WARNING = 1
    ERROR = 2

class Notification(graphene.ObjectType):
    title = graphene.String(required=True)
    kind = graphene.NonNull(NotificationKind)

class ContractStatus(graphene.Enum):
    INACTIVE = 0
    ACTIVE = 1

class Contract(graphene.ObjectType):
    amount = graphene.Float(required=True)
    status = graphene.NonNull(ContractStatus)
    original_id = graphene.String(required=True)

class ContractOuter(graphene.Enum):
    ALL = 0
    PROFILE = 1

class ContractFilter(graphene.InputObjectType):
    outer = graphene.NonNull(ContractOuter) #

class PaginatedContracts(graphene.ObjectType):
    total_count = graphene.Int(required=True)
    nodes = graphene.List(graphene.NonNull(Contract))

class Viewer(graphene.ObjectType):
    name = graphene.String(required=False)
    avatar = graphene.Field(Image)
    notifications = graphene.List(graphene.NonNull(Notification))

    def resolve_notifications(root, _info):
        _notifications = (
            {"kind": NotificationKind.ERROR, "title": "i am a error", "owner_id": "1" },
            {"kind": NotificationKind.WARNING, "title": "i am a warning", "owner_id": "2" },
            {"kind": NotificationKind.INFO, "title": "i am a info", "owner_id": "1" },
        )

        return [n for n in _notifications if n['owner_id'] == root['id']]

class Query(graphene.ObjectType):
    viewer = graphene.Field(Viewer)
    contracts = graphene.Field(
        PaginatedContracts,
        filter=graphene.Argument(ContractFilter),
        order_by=graphene.Argument(graphene.List(graphene.NonNull(OrderBy)))
    )

    def resolve_viewer(_root, _info):
        avatar = { "href": "https://avatar.com/1.jpg", "title": "Samuel Avatar" }

        return {
            "id": "2",
            "name": "Samuel Pythonando",
            "avatar": avatar
        }

    def resolve_contracts(_root, _info, **kwargs):
        _contracts = (
            {"owner_id": "1", "amount": 11.11, "original_id": "111", "status": ContractStatus.ACTIVE },
            {"owner_id": "2", "amount": 22.22, "original_id": "222", "status": ContractStatus.INACTIVE },
            {"owner_id": "2", "amount": 33.33, "original_id": "333", "status": ContractStatus.INACTIVE },
            {"owner_id": "1", "amount": 44.44, "original_id": "444", "status": ContractStatus.ACTIVE },
        )

        if kwargs['filter']['outer'] == ContractOuter.ALL:
            return {
                "total_count": 0,
                "nodes": _contracts,
            }

        _filtered = [c for c in _contracts if c['owner_id'] == "1"]

        return {
            "total_count": len(_filtered),
            "nodes": _filtered,
        }


schema = graphene.Schema(query=Query)
