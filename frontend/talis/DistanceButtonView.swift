//
//  DistanceButtonView.swift
//  talis
//
//  Created by Jordan Jimenez on 5/3/24.
//

import SwiftUI

struct DistanceButtonView: View {
    var options: [String] = ["\u{1F9CD}\u{1F3DF}", "\u{1F9CD}   \u{1F3DF}","\u{1F9CD}        \u{1F3DF}"]

    @State var currentIndex: Int = 0
    @Binding var distance: String
    var body: some View {
        Button(action:{
            
            currentIndex = (currentIndex + 1) % options.count
            
            let currentDistance = options[currentIndex]
            switch currentDistance{
            case "\u{1F9CD}\u{1F3DF}":
                distance = "1"
            case "\u{1F9CD}   \u{1F3DF}":
                distance = "2"
            case "\u{1F9CD}        \u{1F3DF}":
                distance = "3"
            default:
                distance = "1"
                
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

