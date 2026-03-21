import ActivityKit
import SwiftUI
import WidgetKit

struct DeliveryLiveActivity: Widget {
    var body: some WidgetConfiguration {
        ActivityConfiguration(for: DeliveryAttributes.self) { context in
            // Lock Screen / Banner UI
            LockScreenView(context: context)
        } dynamicIsland: { context in
            DynamicIsland {
                // Expanded Dynamic Island
                DynamicIslandExpandedRegion(.leading) {
                    VStack(alignment: .leading, spacing: 4) {
                        Label(context.attributes.restaurantName, systemImage: "storefront.fill")
                            .font(.caption2.bold())
                        Text(context.attributes.orderNumber)
                            .font(.caption2)
                            .foregroundStyle(.secondary)
                    }
                }

                DynamicIslandExpandedRegion(.trailing) {
                    VStack(alignment: .trailing, spacing: 4) {
                        Text("\(context.state.estimatedMinutes) min")
                            .font(.title3.bold())
                            .foregroundStyle(statusColor(for: context.state.status))
                        Text("ETA")
                            .font(.caption2)
                            .foregroundStyle(.secondary)
                    }
                }

                DynamicIslandExpandedRegion(.center) {
                    VStack(spacing: 6) {
                        Image(systemName: context.state.status.icon)
                            .font(.title2)
                            .foregroundStyle(statusColor(for: context.state.status))

                        Text(context.state.status.label)
                            .font(.caption.bold())
                            .foregroundStyle(statusColor(for: context.state.status))
                    }
                }

                DynamicIslandExpandedRegion(.bottom) {
                    ProgressView(value: progress(for: context.state.status))
                        .tint(statusColor(for: context.state.status))
                        .padding(.horizontal)

                    HStack {
                        Image(systemName: "person.circle.fill")
                            .foregroundStyle(.secondary)
                        Text(context.state.driverName)
                            .font(.caption)
                            .foregroundStyle(.secondary)
                        Spacer()
                        Text("Driver")
                            .font(.caption2)
                            .foregroundStyle(.tertiary)
                    }
                    .padding(.horizontal)
                    .padding(.top, 4)
                }
            } compactLeading: {
                Image(systemName: context.state.status.icon)
                    .foregroundStyle(statusColor(for: context.state.status))
            } compactTrailing: {
                Text("\(context.state.estimatedMinutes)m")
                    .font(.caption2.bold())
                    .foregroundStyle(statusColor(for: context.state.status))
            } minimal: {
                Image(systemName: context.state.status.icon)
                    .foregroundStyle(statusColor(for: context.state.status))
            }
        }
    }

    private func statusColor(for status: DeliveryStatus) -> Color {
        switch status {
        case .preparing: return .orange
        case .pickedUp: return .blue
        case .onTheWay: return .purple
        case .arriving, .delivered: return .green
        }
    }

    private func progress(for status: DeliveryStatus) -> Double {
        switch status {
        case .preparing: return 0.2
        case .pickedUp: return 0.4
        case .onTheWay: return 0.65
        case .arriving: return 0.9
        case .delivered: return 1.0
        }
    }
}

struct LockScreenView: View {
    let context: ActivityViewContext<DeliveryAttributes>

    var body: some View {
        VStack(spacing: 12) {
            HStack {
                VStack(alignment: .leading, spacing: 4) {
                    Text(context.attributes.restaurantName)
                        .font(.headline)
                    Text(context.attributes.orderNumber)
                        .font(.subheadline)
                        .foregroundStyle(.secondary)
                }

                Spacer()

                VStack(alignment: .trailing, spacing: 4) {
                    Text("\(context.state.estimatedMinutes) min")
                        .font(.title2.bold())
                    Text("remaining")
                        .font(.caption)
                        .foregroundStyle(.secondary)
                }
            }

            HStack(spacing: 12) {
                Image(systemName: context.state.status.icon)
                    .font(.title3)
                    .foregroundStyle(statusColor)

                VStack(alignment: .leading, spacing: 2) {
                    Text(context.state.status.label)
                        .font(.subheadline.bold())
                    Text("Driver: \(context.state.driverName)")
                        .font(.caption)
                        .foregroundStyle(.secondary)
                }

                Spacer()

                ProgressView(value: progress)
                    .tint(statusColor)
                    .frame(width: 80)
            }
        }
        .padding()
        .activityBackgroundTint(.black.opacity(0.8))
        .activitySystemActionForegroundColor(.white)
    }

    private var statusColor: Color {
        switch context.state.status {
        case .preparing: return .orange
        case .pickedUp: return .blue
        case .onTheWay: return .purple
        case .arriving, .delivered: return .green
        }
    }

    private var progress: Double {
        switch context.state.status {
        case .preparing: return 0.2
        case .pickedUp: return 0.4
        case .onTheWay: return 0.65
        case .arriving: return 0.9
        case .delivered: return 1.0
        }
    }
}
