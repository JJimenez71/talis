//
//  ActivityInfoView.swift
//  talis
//
//  Created by Jordan Jimenez on 5/7/24.
//

import SwiftUI

struct ActivityInfoView: View {
    var activity: Activity
    var body: some View {
        GeometryReader{ geometry in
            if activity.Name != nil{
                VStack{
                    Text(activity.Name!).font(.title).foregroundStyle(.black)
                    AsyncImage(url: URL(string:activity.image!)){ image in
                        image.image?.resizable()
                    }.frame(width: 256, height: 256)
                    Text(activity.Address!).font(.system(size: geometry.size.width * 0.05))
                    Text(activity.phone!).font(.title2)
                    Text(activity.rating!)
                    Text(activity.website!).font(.title3)
                }.frame(width: geometry.size.width)
            }
        }
    }
}

struct ActivityInfoViewPreview: PreviewProvider{
    @State static var act: Activity = Activity(Address: "1235 E 370 S Payson Utah 84651 12345234", Name: "My old house", image: "https://www.yelp.com/biz/bistro-provenance-provo-2?adjust_creative=kPNStrI5K6gGO2ZcCDJxvw\\u0026utm_campaign=yelp_api_v3\\u0026utm_medium=api_v3_business_search\\u0026utm_source=kPNStrI5K6gGO2ZcCDJxvw", phone: "801-465-9698", rating: "5.0", website: "http://myoldhome.com")
    static var previews: some View{
        ActivityInfoView(activity: act)
    }
}

