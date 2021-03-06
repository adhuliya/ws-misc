; A hello world program in llvm ir from https://llvm.org/docs/LangRef.html
; Compile using: `clang helloworld.ll`

; Declare the string constant as a global constant.
@..str = private unnamed_addr constant [14 x i8] c"Hello, World\0A\00"

; External declaration of the puts function
declare i32 @puts(i8* nocapture) nounwind

; Definition of main function
define i32 @main() {   ; i32()*
  ; Convert [14 x i8]* to i8*...
  %cast210 = getelementptr [14 x i8], [14 x i8]* @..str, i64 0, i64 0

  ; Call puts function to write out the string to stdout.
  call i32 @puts(i8* %cast210)
  ret i32 0
}

; Named metadata
!0 = !{i32 42, null, !"string"}
!foo = !{!0}
