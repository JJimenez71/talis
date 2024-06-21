//
//  PriceButtonView.swift
//  talis
//
//  Created by Jordan Jimenez on 4/28/24.
//

import SwiftUI

struct PriceButtonView: View {
    var options = ["$", "$$", "$$$", "$$$$"]
    @State var currentIndex: Int = 0
    @Binding var expense: String
    var body: some View {
        Button(action:{
            
            currentIndex = (currentIndex + 1) % options.count
            
            let currentExpense = options[currentIndex]
            switch currentExpense{
            case "$":
                expense = "1"
            case "$$":
                expense = "2"
            case "$$$":
                expense = "3"
            case "$$$$":
                expense = "4"
            default:
                expense = "1"
                
            }
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


struct PriceButtonPreview: PreviewProvider {
    @State static private var expense: String = "1"
    static var previews: some View{
        PriceButtonView(expense: $expense)
    }
}

