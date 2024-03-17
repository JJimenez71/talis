//
//  LocationManager.swift
//  diceactivity
//
//  Created by Jordan Jimenez on 3/10/24.
//

import Foundation
import CoreLocation

class LocationManager: NSObject, CLLocationManagerDelegate, ObservableObject{
    var locationManager = CLLocationManager()
    @Published var authorizationStatus: CLAuthorizationStatus?
    
    override init(){
        super.init()
        locationManager.delegate = self
    }
    
    func locationManagerDidChangeAuthorization(_ manager: CLLocationManager) {
        switch manager.authorizationStatus{
        case .authorizedWhenInUse:
            print("Authorized")
            authorizationStatus = .authorizedWhenInUse
            locationManager.requestLocation()
            break
        case .restricted:
            print("Restricted or Denied")
            authorizationStatus = .restricted
            break
        case .denied:
            print("Denied")
            authorizationStatus = .denied
        case .notDetermined:
            print("Not determined")
            authorizationStatus = .notDetermined
            locationManager.requestWhenInUseAuthorization()
            break
        default:
            print("broken")
            break
        }
    }
    
    func locationManager(_ manager: CLLocationManager, didUpdateLocations locations: [CLLocation]) {
        print("Updated Location")
    }
    
    func locationManager(_ manager: CLLocationManager, didFailWithError error: any Error) {
        print("Error: \(error.localizedDescription)")
    }

}
