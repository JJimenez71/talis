//
//  RollButtonView.swift
//  diceactivity
//
//  Created by Jordan Jimenez on 3/10/24.
//

import SwiftUI

struct RollButtonView: View {
    var body: some View {
        Button("Roll"){
            print("We rolled")
        }
        .padding()
        .font(/*@START_MENU_TOKEN@*/.title/*@END_MENU_TOKEN@*/)
        .foregroundColor(.white)
        .background(Color(red: 252/255, green:191/255, blue: 73/255))
        .clipShape(Capsule())
        
    }
}

#Preview {
    RollButtonView()
}
