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
    
    // Options for filter buttons
    @State var distanceOptions: [String] = ["\u{1F9CD}\u{1F3DF}", "\u{1F9CD}   \u{1F3DF}","\u{1F9CD}        \u{1F3DF}"]
    @State var priceOptions: [String] = ["$", "$$", "$$$", "$$$$"]
    
    var body: some View {
        Group {
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
            VStack{
                RollButtonView()
                HStack{
                    RotateButtonView(options: $distanceOptions)
                    RotateButtonView(options: $priceOptions)
                }
            }
            .padding()
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Color.gray)
    }
}

#Preview {
    ContentView()
}
