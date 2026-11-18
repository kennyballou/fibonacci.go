(define-module (manifest)
  #:use-module (guix profiles)
  #:use-module (gnu packages base)
  #:use-module (gnu packages commencement)
  #:use-module (gnu packages golang)
  #:export (%project-manifest))

(define %project-manifest
  (packages->manifest
   (list go
         gcc-toolchain
         gnu-make)))

%project-manifest
