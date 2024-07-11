//
//  RollButtonView.swift
//  diceactivity
//
//  Created by Jordan Jimenez on 3/10/24.
//

import SwiftUI


struct RollButtonView: View {
    
    @Binding var activity: Activity
    @Binding var expense: String
    @Binding var distance: String
    var latitude: Double
    var longitude: Double
    var body: some View {
        GeometryReader{ geometry in
            VStack{
                Spacer()
                HStack{
                    Spacer()
                    Button(action:{
                        Task{
                            let url = URL(string: "http://10.1.11.195:8080/roll?distance=\(distance)&expense=\(expense)&latitude=\(latitude)&longitude=\(longitude)")!
                            
                            RequestController.shared.getJSON(url: url, completion: {(result: Result<Activity, Error>) -> Void in
                                switch result{
                                case .success(let jsonData):
                                    activity = jsonData
                                    activity.rating = String(format: "%.2f", Double(activity.rating!) ?? 0.0) // Format our rating
                                    print(jsonData)
                                case .failure(let err):
                                    print(err)
                                }
                            })
                        }
                    }){
                        Text("Roll").foregroundStyle(Color(red: 91.0, green:159.0, blue: 92.0)).frame(minWidth: 0, maxWidth: .infinity)
                    }
                    .buttonStyle(TalisButtonStyle())
                    .frame(minWidth: geometry.size.width * 0.8)
                    Spacer()
                }
            }
        }
        
    }
}


struct RollButtonPreview: PreviewProvider{
    @State static var act: Activity = Activity(address: "1235 E 370 S Payson Utah 84651 12345234", name: "My old house", image: "https://www.yelp.com/biz/bistro-provenance-provo-2?adjust_creative=kPNStrI5K6gGO2ZcCDJxvw\\u0026utm_campaign=yelp_api_v3\\u0026utm_medium=api_v3_business_search\\u0026utm_source=kPNStrI5K6gGO2ZcCDJxvw", phone: "801-465-9698", rating: "5.0", website: "http://myoldhome.com")
    @State static var expense: String = "1"
    @State static var distance: String = "1"
    static var previews: some View{
        RollButtonView(activity: $act, expense: $expense, distance: $distance, latitude: 10.0, longitude: 10.0)
    }
}
