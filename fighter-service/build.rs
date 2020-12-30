fn main() {
    let iface_files = &["../protos/greeter.proto"];
    let dirs = &["../"];

    tonic_build::configure()
        .build_client(true)
        .out_dir(&"./src")
        .compile(iface_files, dirs)
        .unwrap_or_else(|e| panic!("protobuf compilation failed: {}", e));

    // recompile protobufs only if any of the proto files changes.
    for file in iface_files {
        println!("cargo:rerun-if-changed={}", file);
    }
}
