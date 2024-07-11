//
//  ContentView.swift
//  diceactivity
//
//  Created by Jordan Jimenez on 3/10/24.
//

import SwiftUI
import CoreLocation


struct Activity: Codable{
    var address: String? = nil
    var name: String? = nil
    var image: String? = nil
    var phone: String? = nil
    var rating: String? = nil
    var website: String? = nil
}


struct ContentView: View {
    // Grab our users location
    @StateObject var locManager = LocationManager()
    
    // Options for filter buttons
    @State private var expense: String = "1"
    @State private var distance: String = "1"
    @State private var activity: Activity = Activity()
    
    private var lat: Double{
        if let l = locManager.latitude{
            return l
        }
        return 0.0
    }
    private var long: Double{
        if let l = locManager.longitude{
            return l
        }
        return 0.0
    }
    
    var body: some View {
        Group {
            Spacer()
            VStack{
                ActivityInfoView(activity: activity)
                RollButtonView(activity: $activity, expense: $expense, distance: $distance, latitude: lat, longitude: long)
                HStack{
                    PriceButtonView(expense: $expense)
                    DistanceButtonView(distance: $distance)
                }
            }
            .padding()
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Color(red: 239.0, green: 222.0, blue: 177.0))
    }
}

#Preview {
    ContentView()
//    Text("For the love of god please learn how to preview")
}


