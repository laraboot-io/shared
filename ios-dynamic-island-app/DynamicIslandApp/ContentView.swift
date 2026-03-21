import SwiftUI

struct ContentView: View {
    @StateObject private var deliveryManager = DeliveryManager()

    var body: some View {
        NavigationStack {
            VStack(spacing: 24) {
                Spacer()

                // Status card
                VStack(spacing: 16) {
                    Image(systemName: deliveryManager.currentStatus.icon)
                        .font(.system(size: 60))
                        .foregroundStyle(statusColor)
                        .contentTransition(.symbolEffect(.replace))

                    Text(deliveryManager.currentStatus.label)
                        .font(.title2.bold())

                    if deliveryManager.isActivityActive && deliveryManager.currentStatus != .delivered {
                        Text("Estimated: \(deliveryManager.estimatedMinutes) min")
                            .font(.subheadline)
                            .foregroundStyle(.secondary)
                    }
                }
                .padding(32)
                .frame(maxWidth: .infinity)
                .background(.ultraThinMaterial, in: RoundedRectangle(cornerRadius: 20))

                Spacer()

                // Controls
                VStack(spacing: 12) {
                    if deliveryManager.isActivityActive {
                        Button {
                            deliveryManager.advanceStatus()
                        } label: {
                            Label("Next Stage", systemImage: "forward.fill")
                                .frame(maxWidth: .infinity)
                        }
                        .buttonStyle(.borderedProminent)
                        .disabled(deliveryManager.currentStatus == .delivered)

                        Button(role: .destructive) {
                            deliveryManager.stopDelivery()
                        } label: {
                            Label("Cancel Delivery", systemImage: "xmark.circle.fill")
                                .frame(maxWidth: .infinity)
                        }
                        .buttonStyle(.bordered)
                    } else {
                        Button {
                            deliveryManager.startDelivery()
                        } label: {
                            Label("Start Delivery", systemImage: "bicycle")
                                .frame(maxWidth: .infinity)
                        }
                        .buttonStyle(.borderedProminent)
                        .tint(.orange)
                    }
                }
                .controlSize(.large)
                .padding(.horizontal)

                Spacer()
            }
            .padding()
            .navigationTitle("Dynamic Island Demo")
        }
    }

    private var statusColor: Color {
        switch deliveryManager.currentStatus {
        case .preparing: return .orange
        case .pickedUp: return .blue
        case .onTheWay: return .purple
        case .arriving, .delivered: return .green
        }
    }
}

#Preview {
    ContentView()
}
