//
//  ActivityNameView.swift
//  talis
//
//  Created by Jordan Jimenez on 7/9/24.
//

import SwiftUI

struct ActivityNameView: View {
    var name: String
    var body: some View {
        Text(name).font(.title).fontWeight(.bold).foregroundStyle(Color(red: 145/255, green: 179/255, blue: 127/255))
    }
}

struct ActivityNameViewPreview: PreviewProvider{
    static var name: String = "Bartolo's"
    static var previews: some View{
        ActivityNameView(name: name)
    }
}
