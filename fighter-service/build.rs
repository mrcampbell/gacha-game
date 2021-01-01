fn main() {
    let iface_files = &["greeter.proto","fighter.proto"];
    let dirs = &["../protos"];

    tonic_build::configure()
        .build_client(true)
        .compile(iface_files, dirs)
        .unwrap_or_else(|e| panic!("protobuf compilation failed: {}", e));

    // recompile protobufs only if any of the proto files changes.
    for file in iface_files {
        println!("cargo:rerun-if-changed={}", file);
    }
}

// fn main() -> Result<(), Box<dyn std::error::Error>> {
//     tonic_build::compile_protos("../protos/greeter.proto")?;
//     Ok(())
// }