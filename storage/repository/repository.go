package repository

import (
	"fmt"

	"fyne.io/fyne"
)

// Repository represents a storage repository, which is a set of methods which
// implement specific functions on a URI. Repositories are registered to handle
// specific URI schemes, and the higher-level functions that operate on URIs
// internally look up an appropriate method from the relevant Repository.
//
// The repository interface includes only methods which must be implemented at
// a minimum. Without implementing all of the methods in this interface, a URI
// would not be usable in a useful way. Additional functionality can be exposed
// by using interfaces which extend Repository.
//
// Repositories are registered to handle a specific URI scheme (or schemes)
// using the Register() method. When a higher-level URI function such as
// storage.Copy() is called, the storage package will internally look up
// the repository associated with the scheme of the URI, then it will use
// a type assertion to check if the repository implements CopyableRepository.
// If so, the Copy() function will be run from the repository, otherwise
// storage.Copy() will return NotSupportedError. This works similarly for
// all other methods in repository-related interfaces.
//
// Note that a repository can be registered for multiple URI schemes. In such
// cases, the repository must internally select and implement the correct
// behavior for each URI scheme.
//
// NOTE: most developers who use Fyne should *not* generally attempt to
// call repository methods directly. You should use the methods in the storage
// package, which will automatically detect the scheme of a URI and call into
// the appropriate repository.
//
// Since: 2.0.0
type Repository interface {

	// Exists will be used to implement calls to storage.Exists() for the
	// registered scheme of this repository.
	//
	// Since 2.0.0
	Exists(u fyne.URI) (bool, error)

	// ReaderFrom will be used to implement calls to storage.ReaderFrom()
	// for the registered scheme of this repository.
	//
	// Since 2.0.0
	ReaderFrom(u fyne.URI) (fyne.URIReadCloser, error)

	// CanRead will be used to implement calls to storage.CanRead() for the
	// registered scheme of this repository.
	//
	// Since 2.0.0
	CanRead(u fyne.URI) (bool, error)

	// Destroy is called when the repository is un-registered from a given
	// URI scheme.
	//
	// Since 2.0.0
	Destroy()
}

// WriteableRepository is an extension of the Repository interface which also
// supports obtaining a writer for URIs of the scheme it is registered to.
//
// Since 2.0.0
type WriteableRepository interface {
	Repository

	// Writer will be used to implement calls to storage.WriterTo() for
	// the registered scheme of this repository.
	//
	// Since 2.0.0
	Writer(u fyne.URI) (fyne.URIWriteCloser, error)

	// CanWrite will be used to implement calls to storage.CanWrite() for the
	// registered scheme of this repository.
	//
	// Since 2.0.0
	CanWrite(u fyne.URI) (bool, error)

	// Delete will be used to implement calls to storage.Delete() for the
	// registered scheme of this repository.
	//
	// Since 2.0.0
	Delete(u fyne.URI) error
}

// ListableRepository is an extension of the Repository interface which also
// supports obtaining directory listings (generally analogous to a directory
// listing) for URIs of the scheme it is registered to.
//
// Since 2.0.0
type ListableRepository interface {
	Repository

	// CanList will be used to implement calls to storage.Listable() for
	// the registered scheme of this repository.
	//
	// Since 2.0.0
	CanList(u fyne.URI) (bool, error)

	// List will be used to implement calls to storage.List() for the
	// registered scheme of this repository.
	//
	// Since 2.0.0
	List(u fyne.URI) ([]fyne.URI, error)
}

// HierarchicalRepository is an extension of the Repository interface which
// also supports determining the parent and child items of a URI.
//
// Since 2.0.0
type HierarchicalRepository interface {
	Repository

	// Parent will be used to implement calls to storage.Parent() for the
	// registered scheme of this repository.
	//
	// A generic implementation is provided in GenericParent(), which
	// is based on the RFC3986 definition of a URI parent.
	//
	// Since 2.0.0
	Parent(fyne.URI) (fyne.URI, error)

	// Child will be used to implement calls to storage.Child() for
	// the registered scheme of this repository.
	//
	// A generic implementation is provided in GenericParent(), which
	// is based on RFC3986.
	//
	// Since 2.0.0
	Child(fyne.URI) (fyne.URI, error)
}

// CopyableRepository is an extension of the Repository interface which also
// supports copying referenced resources from one URI to another.
type CopyableRepository interface {
	Repository

	// Copy will be used to implement calls to storage.Copy() for the
	// registered scheme of this repository.
	//
	// A generic implementation is provided by GenericCopy().
	//
	// NOTE: the first parameter is the source, the second is the
	// destination.
	//
	// NOTE: if storage.Copy() is given two URIs of different schemes, it
	// is possible that only the source URI will be of the type this
	// repository is registered to handle. In such cases, implementations
	// are suggested to fail-over to GenericCopy().
	//
	// Since 2.0.0
	Copy(fyne.URI, fyne.URI) error
}

// MovableRepository is an extension of the Repository interface which also
// supports moving referenced resources from one URI to another.
type MovableRepository interface {
	Repository

	// Move will be used to implement calls to storage.Move() for the
	// registered scheme of this repository.
	//
	// A generic implementation is provided by GenericMove().
	//
	// NOTE: the first parameter is the source, the second is the
	// destination.
	//
	// NOTE: if storage.Move() is given two URIs of different schemes, it
	// is possible that only the source URI will be of the type this
	// repository is registered to handle. In such cases, implementations
	// are suggested to fail-over to GenericMove().
	//
	// Since 2.0.0
	Move(fyne.URI, fyne.URI) error
}

// Register registers a storage repository so that operations on URIs of the
// registered scheme will use methods implemented by the relevant repository
// implementation.
//
// Since 2.0.0
func Register(scheme string, repository Repository) {
}

// RegisteredRepository returns the Repository instance which is registered to
// handle URIs of the given scheme.
//
// NOTE: this function is intended to be used specifically by the storage
// package. It generally should not be used outside of the fyne package -
// instead you should use the methods in the storage package.
func RegisteredRepository(u fyne.URI) (Repository, error) {
	return nil, fmt.Errorf("TODO")
}