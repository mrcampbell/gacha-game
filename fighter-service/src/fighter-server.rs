use tonic::{transport::Server, Request, Response, Status};

use fighter::fighter_service_server::{FighterService, FighterServiceServer};
use fighter::{Element, Unit, UnitByIdRequest, UnitByIdResponse, UnitType};

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
    println!("Received request from: {:?}", request);

    let unit = Unit {
      id: format!("Hello {}!", request.into_inner().id).into(),
      // element: Element::Fire,
      element: 1,
      r#type: 2,
      // r#type: UnitType::Tank,
      name: "Falcano".to_string(),
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
