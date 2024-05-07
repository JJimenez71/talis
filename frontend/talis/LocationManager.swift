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
    var latitude: Double?
    var longitude: Double?
    override init(){
        super.init()
        locationManager.delegate = self
    }
    
    func locationManagerDidChangeAuthorization(_ manager: CLLocationManager) {
        switch manager.authorizationStatus{
        case .authorizedWhenInUse:
            authorizationStatus = .authorizedWhenInUse
            latitude = manager.location?.coordinate.latitude
            longitude = manager.location?.coordinate.longitude
            locationManager.requestLocation()
            break
        case .restricted:
            authorizationStatus = .restricted
            latitude = 0.0
            longitude = 0.0
            break
        case .denied:
            authorizationStatus = .denied
        case .notDetermined:
            authorizationStatus = .notDetermined
            locationManager.requestWhenInUseAuthorization()
            break
        default:
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
