//
//  RollButtonView.swift
//  diceactivity
//
//  Created by Jordan Jimenez on 3/10/24.
//

import SwiftUI


struct Activity: Codable{
    let name: String
    let address: String
    let hours: [String : String]
    let phone: String
    let website: String
    let rating: Int
    let price: String
    let description: String
}

func roll() async throws -> Activity{
    let apiURL = URL(string: "http://localhost:8080/")!
    let (data, _) = try await URLSession.shared.data(from: apiURL)
    let act = try JSONDecoder().decode(Activity.self, from: data)
    return act
}

func roll_test(){
    print("Hello")
}


struct RollButtonView: View {
    var body: some View {
        GeometryReader{ geometry in
            VStack{
                Spacer()
                HStack{
                    Spacer()
                    Button(action:{
                        roll_test()
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
    RollButtonView()
}
