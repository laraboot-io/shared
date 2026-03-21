import ActivityKit
import Foundation

struct DeliveryAttributes: ActivityAttributes {
    public struct ContentState: Codable, Hashable {
        var status: DeliveryStatus
        var estimatedMinutes: Int
        var driverName: String
    }

    var orderNumber: String
    var restaurantName: String
}

enum DeliveryStatus: String, Codable, Hashable {
    case preparing
    case pickedUp
    case onTheWay
    case arriving
    case delivered

    var label: String {
        switch self {
        case .preparing: return "Preparing"
        case .pickedUp: return "Picked Up"
        case .onTheWay: return "On the Way"
        case .arriving: return "Arriving"
        case .delivered: return "Delivered"
        }
    }

    var icon: String {
        switch self {
        case .preparing: return "frying.pan"
        case .pickedUp: return "bag.fill"
        case .onTheWay: return "car.fill"
        case .arriving: return "mappin.circle.fill"
        case .delivered: return "checkmark.circle.fill"
        }
    }

    var color: String {
        switch self {
        case .preparing: return "orange"
        case .pickedUp: return "blue"
        case .onTheWay: return "purple"
        case .arriving: return "green"
        case .delivered: return "green"
        }
    }
}
