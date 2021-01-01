use tonic::{transport::Server, Request, Response, Status};
extern crate fighter_service;
extern crate diesel;
use self::fighter_service::*;
use self::models::*;
use self::diesel::prelude::*;

use fighter::fighter_service_server::{FighterService, FighterServiceServer};
use fighter::{Unit as pbUnit, UnitByIdRequest, UnitByIdResponse};


// Import the generated proto-rust file into a module
pub mod fighter {
  tonic::include_proto!("fighter");
}

// Implement the service skeleton for the "Fighter" service
// defined in the proto
#[derive(Debug, Default)]
pub struct MyFighterServiceServer {}

// Implement the service function(s) defined in the proto
// for the Greeter service (SayHello...)
#[tonic::async_trait]
impl FighterService for MyFighterServiceServer {
  async fn unit_by_id(
    &self,
    request: Request<UnitByIdRequest>,
  ) -> Result<Response<UnitByIdResponse>, Status> {
    use fighter_service::schema::units::dsl::*;

    let req = request.into_inner();
    println!("Received request from: {:?}", req);


    let connection = establish_connection();
    let result = units
        .filter(id.eq(req.id))
        .first::<Unit>(&connection)
        .expect("Error loading units");

    let unit = pbUnit {
      id: result.id,
      // element: Element::Fire,
      element: result.element,
      r#type: result.type_,
      // r#type: UnitType::Tank,
      // name: "Falcano".to_string(),
      name: result.name
    };

    let response = fighter::UnitByIdResponse { unit: Some(unit) };

    Ok(Response::new(response))
  }
}

// Use the tokio runtime to run our server
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
  let addr = "[::1]:5000".parse()?;
  let server = MyFighterServiceServer::default();

  println!("Starting gRPC Server...");
  Server::builder()
    .add_service(FighterServiceServer::new(server))
    .serve(addr)
    .await?;

  Ok(())
}
