//
//  RollButtonView.swift
//  diceactivity
//
//  Created by Jordan Jimenez on 3/10/24.
//

import SwiftUI


//struct Activity: Codable{
//    var Address: String
//    var Name: String
//    var image: String
//    var phone: String
//    var rating: String
//    var website: String
//}

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
                            print("Expense from Roll: \(expense)")
                            print("Distance from Roll: \(distance)")
                            let url = URL(string: "http://10.1.11.155:8080/roll?distance=\(distance)&expense=\(expense)&latitude=\(latitude)&longitude=\(longitude)")!
                            RequestController.shared.getJSON(url: url, completion: {(result: Result<Activity, Error>) -> Void in
                                switch result{
                                case .success(let jsonData):
                                    activity = jsonData
                                    print(jsonData)
                                case .failure(let err):
                                    print(err)
                                }
                            })
                        }
                    }){
                        Text("Roll")
                    }
                    .padding()
                    .frame(minWidth: geometry.size.width * 0.8)
                    .font(/*@START_MENU_TOKEN@*/.title/*@END_MENU_TOKEN@*/)
                    .foregroundColor(.white)
                    .background(Color(red: 252/255, green:191/255, blue: 73/255))
                    .clipShape(Capsule())
                    Spacer()
                }
            }
        }
        
    }
}

#Preview {
    Text("hello")
}
