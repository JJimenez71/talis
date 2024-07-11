//
//  TalisButtonStyle.swift
//  talis
//
//  Created by Jordan Jimenez on 7/5/24.
//

import Foundation
import SwiftUI

struct TalisButtonStyle: ButtonStyle{
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .padding()
            .font(.title)
            .foregroundColor(.white)
            .background(Color(red: 145/225, green: 179/255, blue: 127/255))
            .clipShape(Capsule())
    }
}

