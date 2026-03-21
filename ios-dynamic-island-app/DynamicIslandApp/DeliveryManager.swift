import ActivityKit
import Foundation

@MainActor
class DeliveryManager: ObservableObject {
    @Published var isActivityActive = false
    @Published var currentStatus: DeliveryStatus = .preparing
    @Published var estimatedMinutes: Int = 30

    private var activity: Activity<DeliveryAttributes>?
    private var simulationTask: Task<Void, Never>?

    func startDelivery() {
        let attributes = DeliveryAttributes(
            orderNumber: "#\(Int.random(in: 1000...9999))",
            restaurantName: "Taco Palace"
        )

        let initialState = DeliveryAttributes.ContentState(
            status: .preparing,
            estimatedMinutes: 30,
            driverName: "Carlos"
        )

        let content = ActivityContent(state: initialState, staleDate: nil)

        do {
            activity = try Activity.request(
                attributes: attributes,
                content: content,
                pushType: nil
            )
            isActivityActive = true
            currentStatus = .preparing
            estimatedMinutes = 30
            startSimulation()
        } catch {
            print("Failed to start Live Activity: \(error.localizedDescription)")
        }
    }

    func stopDelivery() {
        simulationTask?.cancel()
        simulationTask = nil

        Task {
            let finalState = DeliveryAttributes.ContentState(
                status: .delivered,
                estimatedMinutes: 0,
                driverName: "Carlos"
            )
            let finalContent = ActivityContent(state: finalState, staleDate: nil)
            await activity?.end(finalContent, dismissalPolicy: .immediate)
            isActivityActive = false
            currentStatus = .preparing
            estimatedMinutes = 30
        }
    }

    func advanceStatus() {
        let nextStatus: DeliveryStatus
        let nextMinutes: Int

        switch currentStatus {
        case .preparing:
            nextStatus = .pickedUp
            nextMinutes = 25
        case .pickedUp:
            nextStatus = .onTheWay
            nextMinutes = 15
        case .onTheWay:
            nextStatus = .arriving
            nextMinutes = 5
        case .arriving:
            nextStatus = .delivered
            nextMinutes = 0
        case .delivered:
            return
        }

        currentStatus = nextStatus
        estimatedMinutes = nextMinutes

        Task {
            let updatedState = DeliveryAttributes.ContentState(
                status: nextStatus,
                estimatedMinutes: nextMinutes,
                driverName: "Carlos"
            )
            let updatedContent = ActivityContent(state: updatedState, staleDate: nil)
            await activity?.update(updatedContent)

            if nextStatus == .delivered {
                try? await Task.sleep(for: .seconds(3))
                stopDelivery()
            }
        }
    }

    private func startSimulation() {
        simulationTask = Task {
            let stages: [(DeliveryStatus, Int, UInt64)] = [
                (.pickedUp, 25, 8),
                (.onTheWay, 15, 10),
                (.arriving, 5, 8),
                (.delivered, 0, 3),
            ]

            for (status, minutes, delay) in stages {
                try? await Task.sleep(for: .seconds(delay))
                if Task.isCancelled { return }

                currentStatus = status
                estimatedMinutes = minutes

                let state = DeliveryAttributes.ContentState(
                    status: status,
                    estimatedMinutes: minutes,
                    driverName: "Carlos"
                )
                let content = ActivityContent(state: state, staleDate: nil)
                await activity?.update(content)

                if status == .delivered {
                    try? await Task.sleep(for: .seconds(3))
                    if Task.isCancelled { return }
                    let finalContent = ActivityContent(state: state, staleDate: nil)
                    await activity?.end(finalContent, dismissalPolicy: .default)
                    isActivityActive = false
                }
            }
        }
    }
}
