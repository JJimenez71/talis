//
//  RequestController.swift
//  talis
//
//  Created by Jordan Jimenez on 4/23/24.
//

import Foundation


class RequestController{
    static let shared = RequestController()
    
    
    func getRequest(url: URL, completion: @escaping (Result<Data, Error>) -> Void){
        
        let requestTask = URLSession.shared.dataTask(with: url){ data, response, error in
            // Return failure if we error out
            if let error = error {
                completion(.failure(error))
                return
            }
            
            // Check our http response (we always get back and HTTPURLResponse class for response)
            guard let requestResponse = response as? HTTPURLResponse, (200...299).contains(requestResponse.statusCode) else{
                completion(.failure(NSError(domain: "http://localhost:8080", code: 0, userInfo: [NSLocalizedDescriptionKey: "Bad response from the server."])))
                return
            }
            
            // Ensure that we actually received data
            if let data = data{
                completion(.success(data))
            } else{
                completion(.failure(NSError(domain: "http://localhost:8080", code: 1, userInfo: [NSLocalizedDescriptionKey: "No data received from the server"])))
            }
        }
        requestTask.resume()
        
    }
    
    func getJSON<T: Decodable>(url: URL, completion: @escaping (Result<T, Error>) -> Void){
        getRequest(url: url){ result in
            switch result{
            case .success(let data):
                do{
                    let jsonData = try JSONDecoder().decode(T.self, from: data)
                    completion(.success(jsonData))
                } catch{
                    completion(.failure(error)) // Uses generic error from do catch
                }
            case .failure(let err):
                completion(.failure(err))
            }
        }
    }
}
