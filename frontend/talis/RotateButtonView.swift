//
//  RotateButtonView.swift
//  talis
//
//  Created by Jordan Jimenez on 3/13/24.
//

import SwiftUI



struct RotateButtonView: View {
    @Binding private var options: [String]
    @Binding private var returnVal: String
    @State private var currentIndex = 0
    var body: some View {
        Button(action:{
            currentIndex = (currentIndex + 1) % options.count
            returnVal = options[currentIndex]
        }){
           Text(options[currentIndex])
                .frame(minWidth: 0, maxWidth: .infinity)
        }
        .padding()
        .font(.title)
        .foregroundColor(.white)
        .background(Color(red: 252/255, green:191/255, blue: 73/255))
        .clipShape(Capsule())
    }
}

#Preview {
    //RotateButtonView()
    Text("I don't know how to preview with Bindings yet...")
}
