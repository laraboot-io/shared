# Dynamic Island App

An iOS app for experimenting with Dynamic Islands using ActivityKit and Live Activities.

## What it does

A delivery tracker demo that showcases all Dynamic Island presentation modes:

- **Compact** — minimal view in the pill (icon + ETA)
- **Expanded** — full view with restaurant, driver, status, and progress bar
- **Minimal** — single icon when another Live Activity is running
- **Lock Screen** — banner widget with full delivery details

## How to use

1. Open in Xcode 15+
2. Run on an iPhone 14 Pro or later (or simulator)
3. Tap **Start Delivery** to launch a Live Activity
4. Watch it cycle through: Preparing → Picked Up → On the Way → Arriving → Delivered
5. Use **Next Stage** to manually advance, or let the auto-simulation run

## Requirements

- iOS 16.1+
- Xcode 15+
- iPhone 14 Pro or later (for Dynamic Island hardware)

## Project Structure

```
DynamicIslandApp/              # Main app target
├── DynamicIslandApp.swift     # App entry point
├── ContentView.swift          # Main UI with controls
├── DeliveryAttributes.swift   # ActivityAttributes model (shared)
└── DeliveryManager.swift      # Live Activity lifecycle manager

DynamicIslandWidgetExtension/  # Widget extension target
├── DynamicIslandWidgetBundle.swift
└── DeliveryLiveActivity.swift # Dynamic Island & Lock Screen UI
```
