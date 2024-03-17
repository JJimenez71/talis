//
//  ContentView.swift
//  diceactivity
//
//  Created by Jordan Jimenez on 3/10/24.
//

import SwiftUI

struct ContentView: View {
    // Grab our users location
    @StateObject var locManager = LocationManager()
    @State var priceOptions: [String] = ["Near"]
    var body: some View {
        VStack {
            switch locManager.locationManager.authorizationStatus{
            case .authorizedWhenInUse:
                Text("Authorized")
            case .restricted, .denied:
                Text("Location services denied")
            case .notDetermined:
                Text("Need to find your location")
            default:
                Text("Waddup")
            }
            Spacer()
            RollButtonView()
            RotateButtonView(options: priceOptions)
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Color.purple)
    }
}

#Preview {
    ContentView()
}
