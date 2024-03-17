//
//  RotateButtonView.swift
//  talis
//
//  Created by Jordan Jimenez on 3/13/24.
//

import SwiftUI



struct RotateButtonView: View {
    @Binding var options: [String]
    var i = 0
    var body: some View {
        VStack{
            Text(options[i])
        }
    }
}

#Preview {
    RotateButtonView()
}
