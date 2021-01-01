import {credentials} from "grpc"
import { ChannelCredentials, ServiceError } from "grpc";
import {FighterServiceClient} from "../proto/fighter_grpc_pb";
import { Element, Unit, UnitByIDRequest, UnitByIDResponse, UnitType } from "../proto/fighter_pb";


export let NewFighterServiceClient = (address: string) => {
  const client = new FighterServiceClient(address, credentials.createInsecure());
  const req = new UnitByIDRequest();
  req.setId("falcano");
  client.unitByID(req, ((err: ServiceError, resp: UnitByIDResponse) => {
    if (err) {
      console.log(err);
      return
    }
    console.log(resp.toObject())
  }));  

}

