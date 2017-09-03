package wl


const DisplayErrorInvalidObject = 0 // server couldn't find object
const DisplayErrorInvalidMethod = 1 // method doesn't exist on the specified interface
const DisplayErrorNoMemory = 2 // server is out of memory

type DisplayListener interface {
    Error(objectID uint32, code uint32, message string)
    DeleteID(id uint32)
}

// The core global object.  This is a special singleton object.  It
// is used for internal Wayland protocol features.
type Display struct {
    ObjectID
    listener DisplayListener
}

func (this *Display) AddListener(listener DisplayListener) {
    this.listener = listener
}

// The sync request asks the server to emit the 'done' event
// on the returned wl_callback object.  Since requests are
// handled in-order and events are delivered in-order, this can
// be used as a barrier to ensure all previous requests and the
// resulting events have been handled.
// 
// The object returned by this request will be destroyed by the
// compositor after the callback is fired and as such the client must not
// attempt to use it after that point.
// 
// The callback_data passed in the callback is the event serial.
func (this *Display) Sync() (*Callback, error) {
    return nil, nil
}

// This request creates a registry object that allows the client
// to list and bind the global objects available from the
// compositor.
func (this *Display) GetRegistry() (*Registry, error) {
    return nil, nil
}



type RegistryListener interface {
    Global(name uint32, iface string, version uint32)
    GlobalRemove(name uint32)
}

// The singleton global registry object.  The server has a number of
// global objects that are available to all clients.  These objects
// typically represent an actual object in the server (for example,
// an input device) or they are singleton objects that provide
// extension functionality.
// 
// When a client creates a registry object, the registry object
// will emit a global event for each global currently in the
// registry.  Globals come and go as a result of device or
// monitor hotplugs, reconfiguration or other events, and the
// registry will send out global and global_remove events to
// keep the client up to date with the changes.  To mark the end
// of the initial burst of events, the client can use the
// wl_display.sync request immediately after calling
// wl_display.get_registry.
// 
// A client can bind to a global object by using the bind
// request.  This creates a client-side handle that lets the object
// emit events to the client and lets the client invoke requests on
// the object.
type Registry struct {
    ObjectID
    listener RegistryListener
}

func (this *Registry) AddListener(listener RegistryListener) {
    this.listener = listener
}

// Binds a new, client-created object to the server using the
// specified name as the identifier.
func (this *Registry) Bind(name uint32) error {
    return nil
}



type CallbackListener interface {
    Done(callbackData uint32)
}

// Clients can handle the 'done' event to get notified when
// the related request is done.
type Callback struct {
    ObjectID
    listener CallbackListener
}

func (this *Callback) AddListener(listener CallbackListener) {
    this.listener = listener
}



type CompositorListener interface {
}

// A compositor.  This object is a singleton global.  The
// compositor is in charge of combining the contents of multiple
// surfaces into one displayable output.
type Compositor struct {
    ObjectID
    listener CompositorListener
}

func (this *Compositor) AddListener(listener CompositorListener) {
    this.listener = listener
}

// Ask the compositor to create a new surface.
func (this *Compositor) CreateSurface() (*Surface, error) {
    return nil, nil
}

// Ask the compositor to create a new region.
func (this *Compositor) CreateRegion() (*Region, error) {
    return nil, nil
}



type ShmPoolListener interface {
}

// The wl_shm_pool object encapsulates a piece of memory shared
// between the compositor and client.  Through the wl_shm_pool
// object, the client can allocate shared memory wl_buffer objects.
// All objects created through the same pool share the same
// underlying mapped memory. Reusing the mapped memory avoids the
// setup/teardown overhead and is useful when interactively resizing
// a surface or for many small buffers.
type ShmPool struct {
    ObjectID
    listener ShmPoolListener
}

func (this *ShmPool) AddListener(listener ShmPoolListener) {
    this.listener = listener
}

// Create a wl_buffer object from the pool.
// 
// The buffer is created offset bytes into the pool and has
// width and height as specified.  The stride argument specifies
// the number of bytes from the beginning of one row to the beginning
// of the next.  The format is the pixel format of the buffer and
// must be one of those advertised through the wl_shm.format event.
// 
// A buffer will keep a reference to the pool it was created from
// so it is valid to destroy the pool immediately after creating
// a buffer from it.
func (this *ShmPool) CreateBuffer(offset int32, width int32, height int32, stride int32, format uint32) (*Buffer, error) {
    return nil, nil
}

// Destroy the shared memory pool.
// 
// The mmapped memory will be released when all
// buffers that have been created from this pool
// are gone.
func (this *ShmPool) Destroy() error {
    return nil
}

// This request will cause the server to remap the backing memory
// for the pool from the file descriptor passed when the pool was
// created, but using the new size.  This request can only be
// used to make the pool bigger.
func (this *ShmPool) Resize(size int32) error {
    return nil
}




const ShmErrorInvalidFormat = 0 // buffer format is not known
const ShmErrorInvalidStride = 1 // invalid size or stride during pool or buffer creation
const ShmErrorInvalidFd = 2 // mmapping the file descriptor failed


const ShmFormatArgb8888 = 0 // 32-bit ARGB format, [31:0] A:R:G:B 8:8:8:8 little endian
const ShmFormatXrgb8888 = 1 // 32-bit RGB format, [31:0] x:R:G:B 8:8:8:8 little endian
const ShmFormatC8 = 0x20203843 // 8-bit color index format, [7:0] C
const ShmFormatRgb332 = 0x38424752 // 8-bit RGB format, [7:0] R:G:B 3:3:2
const ShmFormatBgr233 = 0x38524742 // 8-bit BGR format, [7:0] B:G:R 2:3:3
const ShmFormatXrgb4444 = 0x32315258 // 16-bit xRGB format, [15:0] x:R:G:B 4:4:4:4 little endian
const ShmFormatXbgr4444 = 0x32314258 // 16-bit xBGR format, [15:0] x:B:G:R 4:4:4:4 little endian
const ShmFormatRgbx4444 = 0x32315852 // 16-bit RGBx format, [15:0] R:G:B:x 4:4:4:4 little endian
const ShmFormatBgrx4444 = 0x32315842 // 16-bit BGRx format, [15:0] B:G:R:x 4:4:4:4 little endian
const ShmFormatArgb4444 = 0x32315241 // 16-bit ARGB format, [15:0] A:R:G:B 4:4:4:4 little endian
const ShmFormatAbgr4444 = 0x32314241 // 16-bit ABGR format, [15:0] A:B:G:R 4:4:4:4 little endian
const ShmFormatRgba4444 = 0x32314152 // 16-bit RBGA format, [15:0] R:G:B:A 4:4:4:4 little endian
const ShmFormatBgra4444 = 0x32314142 // 16-bit BGRA format, [15:0] B:G:R:A 4:4:4:4 little endian
const ShmFormatXrgb1555 = 0x35315258 // 16-bit xRGB format, [15:0] x:R:G:B 1:5:5:5 little endian
const ShmFormatXbgr1555 = 0x35314258 // 16-bit xBGR 1555 format, [15:0] x:B:G:R 1:5:5:5 little endian
const ShmFormatRgbx5551 = 0x35315852 // 16-bit RGBx 5551 format, [15:0] R:G:B:x 5:5:5:1 little endian
const ShmFormatBgrx5551 = 0x35315842 // 16-bit BGRx 5551 format, [15:0] B:G:R:x 5:5:5:1 little endian
const ShmFormatArgb1555 = 0x35315241 // 16-bit ARGB 1555 format, [15:0] A:R:G:B 1:5:5:5 little endian
const ShmFormatAbgr1555 = 0x35314241 // 16-bit ABGR 1555 format, [15:0] A:B:G:R 1:5:5:5 little endian
const ShmFormatRgba5551 = 0x35314152 // 16-bit RGBA 5551 format, [15:0] R:G:B:A 5:5:5:1 little endian
const ShmFormatBgra5551 = 0x35314142 // 16-bit BGRA 5551 format, [15:0] B:G:R:A 5:5:5:1 little endian
const ShmFormatRgb565 = 0x36314752 // 16-bit RGB 565 format, [15:0] R:G:B 5:6:5 little endian
const ShmFormatBgr565 = 0x36314742 // 16-bit BGR 565 format, [15:0] B:G:R 5:6:5 little endian
const ShmFormatRgb888 = 0x34324752 // 24-bit RGB format, [23:0] R:G:B little endian
const ShmFormatBgr888 = 0x34324742 // 24-bit BGR format, [23:0] B:G:R little endian
const ShmFormatXbgr8888 = 0x34324258 // 32-bit xBGR format, [31:0] x:B:G:R 8:8:8:8 little endian
const ShmFormatRgbx8888 = 0x34325852 // 32-bit RGBx format, [31:0] R:G:B:x 8:8:8:8 little endian
const ShmFormatBgrx8888 = 0x34325842 // 32-bit BGRx format, [31:0] B:G:R:x 8:8:8:8 little endian
const ShmFormatAbgr8888 = 0x34324241 // 32-bit ABGR format, [31:0] A:B:G:R 8:8:8:8 little endian
const ShmFormatRgba8888 = 0x34324152 // 32-bit RGBA format, [31:0] R:G:B:A 8:8:8:8 little endian
const ShmFormatBgra8888 = 0x34324142 // 32-bit BGRA format, [31:0] B:G:R:A 8:8:8:8 little endian
const ShmFormatXrgb2101010 = 0x30335258 // 32-bit xRGB format, [31:0] x:R:G:B 2:10:10:10 little endian
const ShmFormatXbgr2101010 = 0x30334258 // 32-bit xBGR format, [31:0] x:B:G:R 2:10:10:10 little endian
const ShmFormatRgbx1010102 = 0x30335852 // 32-bit RGBx format, [31:0] R:G:B:x 10:10:10:2 little endian
const ShmFormatBgrx1010102 = 0x30335842 // 32-bit BGRx format, [31:0] B:G:R:x 10:10:10:2 little endian
const ShmFormatArgb2101010 = 0x30335241 // 32-bit ARGB format, [31:0] A:R:G:B 2:10:10:10 little endian
const ShmFormatAbgr2101010 = 0x30334241 // 32-bit ABGR format, [31:0] A:B:G:R 2:10:10:10 little endian
const ShmFormatRgba1010102 = 0x30334152 // 32-bit RGBA format, [31:0] R:G:B:A 10:10:10:2 little endian
const ShmFormatBgra1010102 = 0x30334142 // 32-bit BGRA format, [31:0] B:G:R:A 10:10:10:2 little endian
const ShmFormatYuyv = 0x56595559 // packed YCbCr format, [31:0] Cr0:Y1:Cb0:Y0 8:8:8:8 little endian
const ShmFormatYvyu = 0x55595659 // packed YCbCr format, [31:0] Cb0:Y1:Cr0:Y0 8:8:8:8 little endian
const ShmFormatUyvy = 0x59565955 // packed YCbCr format, [31:0] Y1:Cr0:Y0:Cb0 8:8:8:8 little endian
const ShmFormatVyuy = 0x59555956 // packed YCbCr format, [31:0] Y1:Cb0:Y0:Cr0 8:8:8:8 little endian
const ShmFormatAyuv = 0x56555941 // packed AYCbCr format, [31:0] A:Y:Cb:Cr 8:8:8:8 little endian
const ShmFormatNv12 = 0x3231564e // 2 plane YCbCr Cr:Cb format, 2x2 subsampled Cr:Cb plane
const ShmFormatNv21 = 0x3132564e // 2 plane YCbCr Cb:Cr format, 2x2 subsampled Cb:Cr plane
const ShmFormatNv16 = 0x3631564e // 2 plane YCbCr Cr:Cb format, 2x1 subsampled Cr:Cb plane
const ShmFormatNv61 = 0x3136564e // 2 plane YCbCr Cb:Cr format, 2x1 subsampled Cb:Cr plane
const ShmFormatYuv410 = 0x39565559 // 3 plane YCbCr format, 4x4 subsampled Cb (1) and Cr (2) planes
const ShmFormatYvu410 = 0x39555659 // 3 plane YCbCr format, 4x4 subsampled Cr (1) and Cb (2) planes
const ShmFormatYuv411 = 0x31315559 // 3 plane YCbCr format, 4x1 subsampled Cb (1) and Cr (2) planes
const ShmFormatYvu411 = 0x31315659 // 3 plane YCbCr format, 4x1 subsampled Cr (1) and Cb (2) planes
const ShmFormatYuv420 = 0x32315559 // 3 plane YCbCr format, 2x2 subsampled Cb (1) and Cr (2) planes
const ShmFormatYvu420 = 0x32315659 // 3 plane YCbCr format, 2x2 subsampled Cr (1) and Cb (2) planes
const ShmFormatYuv422 = 0x36315559 // 3 plane YCbCr format, 2x1 subsampled Cb (1) and Cr (2) planes
const ShmFormatYvu422 = 0x36315659 // 3 plane YCbCr format, 2x1 subsampled Cr (1) and Cb (2) planes
const ShmFormatYuv444 = 0x34325559 // 3 plane YCbCr format, non-subsampled Cb (1) and Cr (2) planes
const ShmFormatYvu444 = 0x34325659 // 3 plane YCbCr format, non-subsampled Cr (1) and Cb (2) planes

type ShmListener interface {
    Format(format uint32)
}

// A singleton global object that provides support for shared
// memory.
// 
// Clients can create wl_shm_pool objects using the create_pool
// request.
// 
// At connection setup time, the wl_shm object emits one or more
// format events to inform clients about the valid pixel formats
// that can be used for buffers.
type Shm struct {
    ObjectID
    listener ShmListener
}

func (this *Shm) AddListener(listener ShmListener) {
    this.listener = listener
}

// Create a new wl_shm_pool object.
// 
// The pool can be used to create shared memory based buffer
// objects.  The server will mmap size bytes of the passed file
// descriptor, to use as backing memory for the pool.
func (this *Shm) CreatePool(size int32) (*ShmPool, error) {
    return nil, nil
}



type BufferListener interface {
    Release()
}

// A buffer provides the content for a wl_surface. Buffers are
// created through factory interfaces such as wl_drm, wl_shm or
// similar. It has a width and a height and can be attached to a
// wl_surface, but the mechanism by which a client provides and
// updates the contents is defined by the buffer factory interface.
type Buffer struct {
    ObjectID
    listener BufferListener
}

func (this *Buffer) AddListener(listener BufferListener) {
    this.listener = listener
}

// Destroy a buffer. If and how you need to release the backing
// storage is defined by the buffer factory interface.
// 
// For possible side-effects to a surface, see wl_surface.attach.
func (this *Buffer) Destroy() error {
    return nil
}




const DataOfferErrorInvalidFinish = 0 // finish request was called untimely
const DataOfferErrorInvalidActionMask = 1 // action mask contains invalid values
const DataOfferErrorInvalidAction = 2 // action argument has an invalid value
const DataOfferErrorInvalidOffer = 3 // offer doesn't accept this request

type DataOfferListener interface {
    Offer(mimeType string)
    SourceActions(sourceActions uint32)
    Action(dndAction uint32)
}

// A wl_data_offer represents a piece of data offered for transfer
// by another client (the source client).  It is used by the
// copy-and-paste and drag-and-drop mechanisms.  The offer
// describes the different mime types that the data can be
// converted to and provides the mechanism for transferring the
// data directly from the source client.
type DataOffer struct {
    ObjectID
    listener DataOfferListener
}

func (this *DataOffer) AddListener(listener DataOfferListener) {
    this.listener = listener
}

// Indicate that the client can accept the given mime type, or
// NULL for not accepted.
// 
// For objects of version 2 or older, this request is used by the
// client to give feedback whether the client can receive the given
// mime type, or NULL if none is accepted; the feedback does not
// determine whether the drag-and-drop operation succeeds or not.
// 
// For objects of version 3 or newer, this request determines the
// final result of the drag-and-drop operation. If the end result
// is that no mime types were accepted, the drag-and-drop operation
// will be cancelled and the corresponding drag source will receive
// wl_data_source.cancelled. Clients may still use this event in
// conjunction with wl_data_source.action for feedback.
func (this *DataOffer) Accept(serial uint32, mimeType string) error {
    return nil
}

// To transfer the offered data, the client issues this request
// and indicates the mime type it wants to receive.  The transfer
// happens through the passed file descriptor (typically created
// with the pipe system call).  The source client writes the data
// in the mime type representation requested and then closes the
// file descriptor.
// 
// The receiving client reads from the read end of the pipe until
// EOF and then closes its end, at which point the transfer is
// complete.
// 
// This request may happen multiple times for different mime types,
// both before and after wl_data_device.drop. Drag-and-drop destination
// clients may preemptively fetch data or examine it more closely to
// determine acceptance.
func (this *DataOffer) Receive(mimeType string) error {
    return nil
}

// Destroy the data offer.
func (this *DataOffer) Destroy() error {
    return nil
}

// Notifies the compositor that the drag destination successfully
// finished the drag-and-drop operation.
// 
// Upon receiving this request, the compositor will emit
// wl_data_source.dnd_finished on the drag source client.
// 
// It is a client error to perform other requests than
// wl_data_offer.destroy after this one. It is also an error to perform
// this request after a NULL mime type has been set in
// wl_data_offer.accept or no action was received through
// wl_data_offer.action.
func (this *DataOffer) Finish() error {
    return nil
}

// Sets the actions that the destination side client supports for
// this operation. This request may trigger the emission of
// wl_data_source.action and wl_data_offer.action events if the compositor
// needs to change the selected action.
// 
// This request can be called multiple times throughout the
// drag-and-drop operation, typically in response to wl_data_device.enter
// or wl_data_device.motion events.
// 
// This request determines the final result of the drag-and-drop
// operation. If the end result is that no action is accepted,
// the drag source will receive wl_drag_source.cancelled.
// 
// The dnd_actions argument must contain only values expressed in the
// wl_data_device_manager.dnd_actions enum, and the preferred_action
// argument must only contain one of those values set, otherwise it
// will result in a protocol error.
// 
// While managing an "ask" action, the destination drag-and-drop client
// may perform further wl_data_offer.receive requests, and is expected
// to perform one last wl_data_offer.set_actions request with a preferred
// action other than "ask" (and optionally wl_data_offer.accept) before
// requesting wl_data_offer.finish, in order to convey the action selected
// by the user. If the preferred action is not in the
// wl_data_offer.source_actions mask, an error will be raised.
// 
// If the "ask" action is dismissed (e.g. user cancellation), the client
// is expected to perform wl_data_offer.destroy right away.
// 
// This request can only be made on drag-and-drop offers, a protocol error
// will be raised otherwise.
func (this *DataOffer) SetActions(dndActions uint32, preferredAction uint32) error {
    return nil
}




const DataSourceErrorInvalidActionMask = 0 // action mask contains invalid values
const DataSourceErrorInvalidSource = 1 // source doesn't accept this request

type DataSourceListener interface {
    Target(mimeType string)
    Send(mimeType string)
    Cancelled()
    DndDropPerformed()
    DndFinished()
    Action(dndAction uint32)
}

// The wl_data_source object is the source side of a wl_data_offer.
// It is created by the source client in a data transfer and
// provides a way to describe the offered data and a way to respond
// to requests to transfer the data.
type DataSource struct {
    ObjectID
    listener DataSourceListener
}

func (this *DataSource) AddListener(listener DataSourceListener) {
    this.listener = listener
}

// This request adds a mime type to the set of mime types
// advertised to targets.  Can be called several times to offer
// multiple types.
func (this *DataSource) Offer(mimeType string) error {
    return nil
}

// Destroy the data source.
func (this *DataSource) Destroy() error {
    return nil
}

// Sets the actions that the source side client supports for this
// operation. This request may trigger wl_data_source.action and
// wl_data_offer.action events if the compositor needs to change the
// selected action.
// 
// The dnd_actions argument must contain only values expressed in the
// wl_data_device_manager.dnd_actions enum, otherwise it will result
// in a protocol error.
// 
// This request must be made once only, and can only be made on sources
// used in drag-and-drop, so it must be performed before
// wl_data_device.start_drag. Attempting to use the source other than
// for drag-and-drop will raise a protocol error.
func (this *DataSource) SetActions(dndActions uint32) error {
    return nil
}




const DataDeviceErrorRole = 0 // given wl_surface has another role

type DataDeviceListener interface {
    DataOffer()
    Enter(serial uint32, surface uint32, x uint32, y uint32, id uint32)
    Leave()
    Motion(time uint32, x uint32, y uint32)
    Drop()
    Selection(id uint32)
}

// There is one wl_data_device per seat which can be obtained
// from the global wl_data_device_manager singleton.
// 
// A wl_data_device provides access to inter-client data transfer
// mechanisms such as copy-and-paste and drag-and-drop.
type DataDevice struct {
    ObjectID
    listener DataDeviceListener
}

func (this *DataDevice) AddListener(listener DataDeviceListener) {
    this.listener = listener
}

// This request asks the compositor to start a drag-and-drop
// operation on behalf of the client.
// 
// The source argument is the data source that provides the data
// for the eventual data transfer. If source is NULL, enter, leave
// and motion events are sent only to the client that initiated the
// drag and the client is expected to handle the data passing
// internally.
// 
// The origin surface is the surface where the drag originates and
// the client must have an active implicit grab that matches the
// serial.
// 
// The icon surface is an optional (can be NULL) surface that
// provides an icon to be moved around with the cursor.  Initially,
// the top-left corner of the icon surface is placed at the cursor
// hotspot, but subsequent wl_surface.attach request can move the
// relative position. Attach requests must be confirmed with
// wl_surface.commit as usual. The icon surface is given the role of
// a drag-and-drop icon. If the icon surface already has another role,
// it raises a protocol error.
// 
// The current and pending input regions of the icon wl_surface are
// cleared, and wl_surface.set_input_region is ignored until the
// wl_surface is no longer used as the icon surface. When the use
// as an icon ends, the current and pending input regions become
// undefined, and the wl_surface is unmapped.
func (this *DataDevice) StartDrag(source uint32, origin uint32, icon uint32, serial uint32) error {
    return nil
}

// This request asks the compositor to set the selection
// to the data from the source on behalf of the client.
// 
// To unset the selection, set the source to NULL.
func (this *DataDevice) SetSelection(source uint32, serial uint32) error {
    return nil
}

// This request destroys the data device.
func (this *DataDevice) Release() error {
    return nil
}




const DataDeviceManagerDndActionNone = 0 // no action
const DataDeviceManagerDndActionCopy = 1 // copy action
const DataDeviceManagerDndActionMove = 2 // move action
const DataDeviceManagerDndActionAsk = 4 // ask action

type DataDeviceManagerListener interface {
}

// The wl_data_device_manager is a singleton global object that
// provides access to inter-client data transfer mechanisms such as
// copy-and-paste and drag-and-drop.  These mechanisms are tied to
// a wl_seat and this interface lets a client get a wl_data_device
// corresponding to a wl_seat.
// 
// Depending on the version bound, the objects created from the bound
// wl_data_device_manager object will have different requirements for
// functioning properly. See wl_data_source.set_actions,
// wl_data_offer.accept and wl_data_offer.finish for details.
type DataDeviceManager struct {
    ObjectID
    listener DataDeviceManagerListener
}

func (this *DataDeviceManager) AddListener(listener DataDeviceManagerListener) {
    this.listener = listener
}

// Create a new data source.
func (this *DataDeviceManager) CreateDataSource() (*DataSource, error) {
    return nil, nil
}

// Create a new data device for a given seat.
func (this *DataDeviceManager) GetDataDevice(seat uint32) (*DataDevice, error) {
    return nil, nil
}




const ShellErrorRole = 0 // given wl_surface has another role

type ShellListener interface {
}

// This interface is implemented by servers that provide
// desktop-style user interfaces.
// 
// It allows clients to associate a wl_shell_surface with
// a basic surface.
type Shell struct {
    ObjectID
    listener ShellListener
}

func (this *Shell) AddListener(listener ShellListener) {
    this.listener = listener
}

// Create a shell surface for an existing surface. This gives
// the wl_surface the role of a shell surface. If the wl_surface
// already has another role, it raises a protocol error.
// 
// Only one shell surface can be associated with a given surface.
func (this *Shell) GetShellSurface(surface uint32) (*ShellSurface, error) {
    return nil, nil
}




const ShellSurfaceResizeNone = 0 // no edge
const ShellSurfaceResizeTop = 1 // top edge
const ShellSurfaceResizeBottom = 2 // bottom edge
const ShellSurfaceResizeLeft = 4 // left edge
const ShellSurfaceResizeTopLeft = 5 // top and left edges
const ShellSurfaceResizeBottomLeft = 6 // bottom and left edges
const ShellSurfaceResizeRight = 8 // right edge
const ShellSurfaceResizeTopRight = 9 // top and right edges
const ShellSurfaceResizeBottomRight = 10 // bottom and right edges


const ShellSurfaceTransientInactive = 0x1 // do not set keyboard focus


const ShellSurfaceFullscreenMethodDefault = 0 // no preference, apply default policy
const ShellSurfaceFullscreenMethodScale = 1 // scale, preserve the surface's aspect ratio and center on output
const ShellSurfaceFullscreenMethodDriver = 2 // switch output mode to the smallest mode that can fit the surface, add black borders to compensate size mismatch
const ShellSurfaceFullscreenMethodFill = 3 // no upscaling, center on output and add black borders to compensate size mismatch

type ShellSurfaceListener interface {
    Ping(serial uint32)
    Configure(edges uint32, width int32, height int32)
    PopupDone()
}

// An interface that may be implemented by a wl_surface, for
// implementations that provide a desktop-style user interface.
// 
// It provides requests to treat surfaces like toplevel, fullscreen
// or popup windows, move, resize or maximize them, associate
// metadata like title and class, etc.
// 
// On the server side the object is automatically destroyed when
// the related wl_surface is destroyed. On the client side,
// wl_shell_surface_destroy() must be called before destroying
// the wl_surface object.
type ShellSurface struct {
    ObjectID
    listener ShellSurfaceListener
}

func (this *ShellSurface) AddListener(listener ShellSurfaceListener) {
    this.listener = listener
}

// A client must respond to a ping event with a pong request or
// the client may be deemed unresponsive.
func (this *ShellSurface) Pong(serial uint32) error {
    return nil
}

// Start a pointer-driven move of the surface.
// 
// This request must be used in response to a button press event.
// The server may ignore move requests depending on the state of
// the surface (e.g. fullscreen or maximized).
func (this *ShellSurface) Move(seat uint32, serial uint32) error {
    return nil
}

// Start a pointer-driven resizing of the surface.
// 
// This request must be used in response to a button press event.
// The server may ignore resize requests depending on the state of
// the surface (e.g. fullscreen or maximized).
func (this *ShellSurface) Resize(seat uint32, serial uint32, edges uint32) error {
    return nil
}

// Map the surface as a toplevel surface.
// 
// A toplevel surface is not fullscreen, maximized or transient.
func (this *ShellSurface) SetToplevel() error {
    return nil
}

// Map the surface relative to an existing surface.
// 
// The x and y arguments specify the location of the upper left
// corner of the surface relative to the upper left corner of the
// parent surface, in surface-local coordinates.
// 
// The flags argument controls details of the transient behaviour.
func (this *ShellSurface) SetTransient(parent uint32, x int32, y int32, flags uint32) error {
    return nil
}

// Map the surface as a fullscreen surface.
// 
// If an output parameter is given then the surface will be made
// fullscreen on that output. If the client does not specify the
// output then the compositor will apply its policy - usually
// choosing the output on which the surface has the biggest surface
// area.
// 
// The client may specify a method to resolve a size conflict
// between the output size and the surface size - this is provided
// through the method parameter.
// 
// The framerate parameter is used only when the method is set
// to "driver", to indicate the preferred framerate. A value of 0
// indicates that the client does not care about framerate.  The
// framerate is specified in mHz, that is framerate of 60000 is 60Hz.
// 
// A method of "scale" or "driver" implies a scaling operation of
// the surface, either via a direct scaling operation or a change of
// the output mode. This will override any kind of output scaling, so
// that mapping a surface with a buffer size equal to the mode can
// fill the screen independent of buffer_scale.
// 
// A method of "fill" means we don't scale up the buffer, however
// any output scale is applied. This means that you may run into
// an edge case where the application maps a buffer with the same
// size of the output mode but buffer_scale 1 (thus making a
// surface larger than the output). In this case it is allowed to
// downscale the results to fit the screen.
// 
// The compositor must reply to this request with a configure event
// with the dimensions for the output on which the surface will
// be made fullscreen.
func (this *ShellSurface) SetFullscreen(method uint32, framerate uint32, output uint32) error {
    return nil
}

// Map the surface as a popup.
// 
// A popup surface is a transient surface with an added pointer
// grab.
// 
// An existing implicit grab will be changed to owner-events mode,
// and the popup grab will continue after the implicit grab ends
// (i.e. releasing the mouse button does not cause the popup to
// be unmapped).
// 
// The popup grab continues until the window is destroyed or a
// mouse button is pressed in any other client's window. A click
// in any of the client's surfaces is reported as normal, however,
// clicks in other clients' surfaces will be discarded and trigger
// the callback.
// 
// The x and y arguments specify the location of the upper left
// corner of the surface relative to the upper left corner of the
// parent surface, in surface-local coordinates.
func (this *ShellSurface) SetPopup(seat uint32, serial uint32, parent uint32, x int32, y int32, flags uint32) error {
    return nil
}

// Map the surface as a maximized surface.
// 
// If an output parameter is given then the surface will be
// maximized on that output. If the client does not specify the
// output then the compositor will apply its policy - usually
// choosing the output on which the surface has the biggest surface
// area.
// 
// The compositor will reply with a configure event telling
// the expected new surface size. The operation is completed
// on the next buffer attach to this surface.
// 
// A maximized surface typically fills the entire output it is
// bound to, except for desktop elements such as panels. This is
// the main difference between a maximized shell surface and a
// fullscreen shell surface.
// 
// The details depend on the compositor implementation.
func (this *ShellSurface) SetMaximized(output uint32) error {
    return nil
}

// Set a short title for the surface.
// 
// This string may be used to identify the surface in a task bar,
// window list, or other user interface elements provided by the
// compositor.
// 
// The string must be encoded in UTF-8.
func (this *ShellSurface) SetTitle(title string) error {
    return nil
}

// Set a class for the surface.
// 
// The surface class identifies the general class of applications
// to which the surface belongs. A common convention is to use the
// file name (or the full path if it is a non-standard location) of
// the application's .desktop file as the class.
func (this *ShellSurface) SetClass(class string) error {
    return nil
}




const SurfaceErrorInvalidScale = 0 // buffer scale value is invalid
const SurfaceErrorInvalidTransform = 1 // buffer transform value is invalid

type SurfaceListener interface {
    Enter(output uint32)
    Leave(output uint32)
}

// A surface is a rectangular area that is displayed on the screen.
// It has a location, size and pixel contents.
// 
// The size of a surface (and relative positions on it) is described
// in surface-local coordinates, which may differ from the buffer
// coordinates of the pixel content, in case a buffer_transform
// or a buffer_scale is used.
// 
// A surface without a "role" is fairly useless: a compositor does
// not know where, when or how to present it. The role is the
// purpose of a wl_surface. Examples of roles are a cursor for a
// pointer (as set by wl_pointer.set_cursor), a drag icon
// (wl_data_device.start_drag), a sub-surface
// (wl_subcompositor.get_subsurface), and a window as defined by a
// shell protocol (e.g. wl_shell.get_shell_surface).
// 
// A surface can have only one role at a time. Initially a
// wl_surface does not have a role. Once a wl_surface is given a
// role, it is set permanently for the whole lifetime of the
// wl_surface object. Giving the current role again is allowed,
// unless explicitly forbidden by the relevant interface
// specification.
// 
// Surface roles are given by requests in other interfaces such as
// wl_pointer.set_cursor. The request should explicitly mention
// that this request gives a role to a wl_surface. Often, this
// request also creates a new protocol object that represents the
// role and adds additional functionality to wl_surface. When a
// client wants to destroy a wl_surface, they must destroy this 'role
// object' before the wl_surface.
// 
// Destroying the role object does not remove the role from the
// wl_surface, but it may stop the wl_surface from "playing the role".
// For instance, if a wl_subsurface object is destroyed, the wl_surface
// it was created for will be unmapped and forget its position and
// z-order. It is allowed to create a wl_subsurface for the same
// wl_surface again, but it is not allowed to use the wl_surface as
// a cursor (cursor is a different role than sub-surface, and role
// switching is not allowed).
type Surface struct {
    ObjectID
    listener SurfaceListener
}

func (this *Surface) AddListener(listener SurfaceListener) {
    this.listener = listener
}

// Deletes the surface and invalidates its object ID.
func (this *Surface) Destroy() error {
    return nil
}

// Set a buffer as the content of this surface.
// 
// The new size of the surface is calculated based on the buffer
// size transformed by the inverse buffer_transform and the
// inverse buffer_scale. This means that the supplied buffer
// must be an integer multiple of the buffer_scale.
// 
// The x and y arguments specify the location of the new pending
// buffer's upper left corner, relative to the current buffer's upper
// left corner, in surface-local coordinates. In other words, the
// x and y, combined with the new surface size define in which
// directions the surface's size changes.
// 
// Surface contents are double-buffered state, see wl_surface.commit.
// 
// The initial surface contents are void; there is no content.
// wl_surface.attach assigns the given wl_buffer as the pending
// wl_buffer. wl_surface.commit makes the pending wl_buffer the new
// surface contents, and the size of the surface becomes the size
// calculated from the wl_buffer, as described above. After commit,
// there is no pending buffer until the next attach.
// 
// Committing a pending wl_buffer allows the compositor to read the
// pixels in the wl_buffer. The compositor may access the pixels at
// any time after the wl_surface.commit request. When the compositor
// will not access the pixels anymore, it will send the
// wl_buffer.release event. Only after receiving wl_buffer.release,
// the client may reuse the wl_buffer. A wl_buffer that has been
// attached and then replaced by another attach instead of committed
// will not receive a release event, and is not used by the
// compositor.
// 
// Destroying the wl_buffer after wl_buffer.release does not change
// the surface contents. However, if the client destroys the
// wl_buffer before receiving the wl_buffer.release event, the surface
// contents become undefined immediately.
// 
// If wl_surface.attach is sent with a NULL wl_buffer, the
// following wl_surface.commit will remove the surface content.
func (this *Surface) Attach(buffer uint32, x int32, y int32) error {
    return nil
}

// This request is used to describe the regions where the pending
// buffer is different from the current surface contents, and where
// the surface therefore needs to be repainted. The compositor
// ignores the parts of the damage that fall outside of the surface.
// 
// Damage is double-buffered state, see wl_surface.commit.
// 
// The damage rectangle is specified in surface-local coordinates,
// where x and y specify the upper left corner of the damage rectangle.
// 
// The initial value for pending damage is empty: no damage.
// wl_surface.damage adds pending damage: the new pending damage
// is the union of old pending damage and the given rectangle.
// 
// wl_surface.commit assigns pending damage as the current damage,
// and clears pending damage. The server will clear the current
// damage as it repaints the surface.
// 
// Alternatively, damage can be posted with wl_surface.damage_buffer
// which uses buffer coordinates instead of surface coordinates,
// and is probably the preferred and intuitive way of doing this.
func (this *Surface) Damage(x int32, y int32, width int32, height int32) error {
    return nil
}

// Request a notification when it is a good time to start drawing a new
// frame, by creating a frame callback. This is useful for throttling
// redrawing operations, and driving animations.
// 
// When a client is animating on a wl_surface, it can use the 'frame'
// request to get notified when it is a good time to draw and commit the
// next frame of animation. If the client commits an update earlier than
// that, it is likely that some updates will not make it to the display,
// and the client is wasting resources by drawing too often.
// 
// The frame request will take effect on the next wl_surface.commit.
// The notification will only be posted for one frame unless
// requested again. For a wl_surface, the notifications are posted in
// the order the frame requests were committed.
// 
// The server must send the notifications so that a client
// will not send excessive updates, while still allowing
// the highest possible update rate for clients that wait for the reply
// before drawing again. The server should give some time for the client
// to draw and commit after sending the frame callback events to let it
// hit the next output refresh.
// 
// A server should avoid signaling the frame callbacks if the
// surface is not visible in any way, e.g. the surface is off-screen,
// or completely obscured by other opaque surfaces.
// 
// The object returned by this request will be destroyed by the
// compositor after the callback is fired and as such the client must not
// attempt to use it after that point.
// 
// The callback_data passed in the callback is the current time, in
// milliseconds, with an undefined base.
func (this *Surface) Frame() (*Callback, error) {
    return nil, nil
}

// This request sets the region of the surface that contains
// opaque content.
// 
// The opaque region is an optimization hint for the compositor
// that lets it optimize the redrawing of content behind opaque
// regions.  Setting an opaque region is not required for correct
// behaviour, but marking transparent content as opaque will result
// in repaint artifacts.
// 
// The opaque region is specified in surface-local coordinates.
// 
// The compositor ignores the parts of the opaque region that fall
// outside of the surface.
// 
// Opaque region is double-buffered state, see wl_surface.commit.
// 
// wl_surface.set_opaque_region changes the pending opaque region.
// wl_surface.commit copies the pending region to the current region.
// Otherwise, the pending and current regions are never changed.
// 
// The initial value for an opaque region is empty. Setting the pending
// opaque region has copy semantics, and the wl_region object can be
// destroyed immediately. A NULL wl_region causes the pending opaque
// region to be set to empty.
func (this *Surface) SetOpaqueRegion(region uint32) error {
    return nil
}

// This request sets the region of the surface that can receive
// pointer and touch events.
// 
// Input events happening outside of this region will try the next
// surface in the server surface stack. The compositor ignores the
// parts of the input region that fall outside of the surface.
// 
// The input region is specified in surface-local coordinates.
// 
// Input region is double-buffered state, see wl_surface.commit.
// 
// wl_surface.set_input_region changes the pending input region.
// wl_surface.commit copies the pending region to the current region.
// Otherwise the pending and current regions are never changed,
// except cursor and icon surfaces are special cases, see
// wl_pointer.set_cursor and wl_data_device.start_drag.
// 
// The initial value for an input region is infinite. That means the
// whole surface will accept input. Setting the pending input region
// has copy semantics, and the wl_region object can be destroyed
// immediately. A NULL wl_region causes the input region to be set
// to infinite.
func (this *Surface) SetInputRegion(region uint32) error {
    return nil
}

// Surface state (input, opaque, and damage regions, attached buffers,
// etc.) is double-buffered. Protocol requests modify the pending state,
// as opposed to the current state in use by the compositor. A commit
// request atomically applies all pending state, replacing the current
// state. After commit, the new pending state is as documented for each
// related request.
// 
// On commit, a pending wl_buffer is applied first, and all other state
// second. This means that all coordinates in double-buffered state are
// relative to the new wl_buffer coming into use, except for
// wl_surface.attach itself. If there is no pending wl_buffer, the
// coordinates are relative to the current surface contents.
// 
// All requests that need a commit to become effective are documented
// to affect double-buffered state.
// 
// Other interfaces may add further double-buffered surface state.
func (this *Surface) Commit() error {
    return nil
}

// This request sets an optional transformation on how the compositor
// interprets the contents of the buffer attached to the surface. The
// accepted values for the transform parameter are the values for
// wl_output.transform.
// 
// Buffer transform is double-buffered state, see wl_surface.commit.
// 
// A newly created surface has its buffer transformation set to normal.
// 
// wl_surface.set_buffer_transform changes the pending buffer
// transformation. wl_surface.commit copies the pending buffer
// transformation to the current one. Otherwise, the pending and current
// values are never changed.
// 
// The purpose of this request is to allow clients to render content
// according to the output transform, thus permitting the compositor to
// use certain optimizations even if the display is rotated. Using
// hardware overlays and scanning out a client buffer for fullscreen
// surfaces are examples of such optimizations. Those optimizations are
// highly dependent on the compositor implementation, so the use of this
// request should be considered on a case-by-case basis.
// 
// Note that if the transform value includes 90 or 270 degree rotation,
// the width of the buffer will become the surface height and the height
// of the buffer will become the surface width.
// 
// If transform is not one of the values from the
// wl_output.transform enum the invalid_transform protocol error
// is raised.
func (this *Surface) SetBufferTransform(transform int32) error {
    return nil
}

// This request sets an optional scaling factor on how the compositor
// interprets the contents of the buffer attached to the window.
// 
// Buffer scale is double-buffered state, see wl_surface.commit.
// 
// A newly created surface has its buffer scale set to 1.
// 
// wl_surface.set_buffer_scale changes the pending buffer scale.
// wl_surface.commit copies the pending buffer scale to the current one.
// Otherwise, the pending and current values are never changed.
// 
// The purpose of this request is to allow clients to supply higher
// resolution buffer data for use on high resolution outputs. It is
// intended that you pick the same buffer scale as the scale of the
// output that the surface is displayed on. This means the compositor
// can avoid scaling when rendering the surface on that output.
// 
// Note that if the scale is larger than 1, then you have to attach
// a buffer that is larger (by a factor of scale in each dimension)
// than the desired surface size.
// 
// If scale is not positive the invalid_scale protocol error is
// raised.
func (this *Surface) SetBufferScale(scale int32) error {
    return nil
}

// This request is used to describe the regions where the pending
// buffer is different from the current surface contents, and where
// the surface therefore needs to be repainted. The compositor
// ignores the parts of the damage that fall outside of the surface.
// 
// Damage is double-buffered state, see wl_surface.commit.
// 
// The damage rectangle is specified in buffer coordinates,
// where x and y specify the upper left corner of the damage rectangle.
// 
// The initial value for pending damage is empty: no damage.
// wl_surface.damage_buffer adds pending damage: the new pending
// damage is the union of old pending damage and the given rectangle.
// 
// wl_surface.commit assigns pending damage as the current damage,
// and clears pending damage. The server will clear the current
// damage as it repaints the surface.
// 
// This request differs from wl_surface.damage in only one way - it
// takes damage in buffer coordinates instead of surface-local
// coordinates. While this generally is more intuitive than surface
// coordinates, it is especially desirable when using wp_viewport
// or when a drawing library (like EGL) is unaware of buffer scale
// and buffer transform.
// 
// Note: Because buffer transformation changes and damage requests may
// be interleaved in the protocol stream, it is impossible to determine
// the actual mapping between surface and buffer damage until
// wl_surface.commit time. Therefore, compositors wishing to take both
// kinds of damage into account will have to accumulate damage from the
// two requests separately and only transform from one to the other
// after receiving the wl_surface.commit.
func (this *Surface) DamageBuffer(x int32, y int32, width int32, height int32) error {
    return nil
}




const SeatCapabilityPointer = 1 // the seat has pointer devices
const SeatCapabilityKeyboard = 2 // the seat has one or more keyboards
const SeatCapabilityTouch = 4 // the seat has touch devices

type SeatListener interface {
    Capabilities(capabilities uint32)
    Name(name string)
}

// A seat is a group of keyboards, pointer and touch devices. This
// object is published as a global during start up, or when such a
// device is hot plugged.  A seat typically has a pointer and
// maintains a keyboard focus and a pointer focus.
type Seat struct {
    ObjectID
    listener SeatListener
}

func (this *Seat) AddListener(listener SeatListener) {
    this.listener = listener
}

// The ID provided will be initialized to the wl_pointer interface
// for this seat.
// 
// This request only takes effect if the seat has the pointer
// capability, or has had the pointer capability in the past.
// It is a protocol violation to issue this request on a seat that has
// never had the pointer capability.
func (this *Seat) GetPointer() (*Pointer, error) {
    return nil, nil
}

// The ID provided will be initialized to the wl_keyboard interface
// for this seat.
// 
// This request only takes effect if the seat has the keyboard
// capability, or has had the keyboard capability in the past.
// It is a protocol violation to issue this request on a seat that has
// never had the keyboard capability.
func (this *Seat) GetKeyboard() (*Keyboard, error) {
    return nil, nil
}

// The ID provided will be initialized to the wl_touch interface
// for this seat.
// 
// This request only takes effect if the seat has the touch
// capability, or has had the touch capability in the past.
// It is a protocol violation to issue this request on a seat that has
// never had the touch capability.
func (this *Seat) GetTouch() (*Touch, error) {
    return nil, nil
}

// Using this request a client can tell the server that it is not going to
// use the seat object anymore.
func (this *Seat) Release() error {
    return nil
}




const PointerErrorRole = 0 // given wl_surface has another role


const PointerButtonStateReleased = 0 // the button is not pressed
const PointerButtonStatePressed = 1 // the button is pressed


const PointerAxisVerticalScroll = 0 // vertical axis
const PointerAxisHorizontalScroll = 1 // horizontal axis


const PointerAxisSourceWheel = 0 // a physical wheel rotation
const PointerAxisSourceFinger = 1 // finger on a touch surface
const PointerAxisSourceContinuous = 2 // continuous coordinate space
const PointerAxisSourceWheelTilt = 3 // a physical wheel tilt

type PointerListener interface {
    Enter(serial uint32, surface uint32, surfaceX uint32, surfaceY uint32)
    Leave(serial uint32, surface uint32)
    Motion(time uint32, surfaceX uint32, surfaceY uint32)
    Button(serial uint32, time uint32, button uint32, state uint32)
    Axis(time uint32, axis uint32, value uint32)
    Frame()
    AxisSource(axisSource uint32)
    AxisStop(time uint32, axis uint32)
    AxisDiscrete(axis uint32, discrete int32)
}

// The wl_pointer interface represents one or more input devices,
// such as mice, which control the pointer location and pointer_focus
// of a seat.
// 
// The wl_pointer interface generates motion, enter and leave
// events for the surfaces that the pointer is located over,
// and button and axis events for button presses, button releases
// and scrolling.
type Pointer struct {
    ObjectID
    listener PointerListener
}

func (this *Pointer) AddListener(listener PointerListener) {
    this.listener = listener
}

// Set the pointer surface, i.e., the surface that contains the
// pointer image (cursor). This request gives the surface the role
// of a cursor. If the surface already has another role, it raises
// a protocol error.
// 
// The cursor actually changes only if the pointer
// focus for this device is one of the requesting client's surfaces
// or the surface parameter is the current pointer surface. If
// there was a previous surface set with this request it is
// replaced. If surface is NULL, the pointer image is hidden.
// 
// The parameters hotspot_x and hotspot_y define the position of
// the pointer surface relative to the pointer location. Its
// top-left corner is always at (x, y) - (hotspot_x, hotspot_y),
// where (x, y) are the coordinates of the pointer location, in
// surface-local coordinates.
// 
// On surface.attach requests to the pointer surface, hotspot_x
// and hotspot_y are decremented by the x and y parameters
// passed to the request. Attach must be confirmed by
// wl_surface.commit as usual.
// 
// The hotspot can also be updated by passing the currently set
// pointer surface to this request with new values for hotspot_x
// and hotspot_y.
// 
// The current and pending input regions of the wl_surface are
// cleared, and wl_surface.set_input_region is ignored until the
// wl_surface is no longer used as the cursor. When the use as a
// cursor ends, the current and pending input regions become
// undefined, and the wl_surface is unmapped.
func (this *Pointer) SetCursor(serial uint32, surface uint32, hotspotX int32, hotspotY int32) error {
    return nil
}

// Using this request a client can tell the server that it is not going to
// use the pointer object anymore.
// 
// This request destroys the pointer proxy object, so clients must not call
// wl_pointer_destroy() after using this request.
func (this *Pointer) Release() error {
    return nil
}




const KeyboardKeymapFormatNoKeymap = 0 // no keymap; client must understand how to interpret the raw keycode
const KeyboardKeymapFormatXkbV1 = 1 // libxkbcommon compatible; to determine the xkb keycode, clients must add 8 to the key event keycode


const KeyboardKeyStateReleased = 0 // key is not pressed
const KeyboardKeyStatePressed = 1 // key is pressed

type KeyboardListener interface {
    Keymap(format uint32, size uint32)
    Enter(serial uint32, surface uint32, keys []byte)
    Leave(serial uint32, surface uint32)
    Key(serial uint32, time uint32, key uint32, state uint32)
    Modifiers(serial uint32, modsDepressed uint32, modsLatched uint32, modsLocked uint32, group uint32)
    RepeatInfo(rate int32, delay int32)
}

// The wl_keyboard interface represents one or more keyboards
// associated with a seat.
type Keyboard struct {
    ObjectID
    listener KeyboardListener
}

func (this *Keyboard) AddListener(listener KeyboardListener) {
    this.listener = listener
}

func (this *Keyboard) Release() error {
    return nil
}



type TouchListener interface {
    Down(serial uint32, time uint32, surface uint32, id int32, x uint32, y uint32)
    Up(serial uint32, time uint32, id int32)
    Motion(time uint32, id int32, x uint32, y uint32)
    Frame()
    Cancel()
    Shape(id int32, major uint32, minor uint32)
    Orientation(id int32, orientation uint32)
}

// The wl_touch interface represents a touchscreen
// associated with a seat.
// 
// Touch interactions can consist of one or more contacts.
// For each contact, a series of events is generated, starting
// with a down event, followed by zero or more motion events,
// and ending with an up event. Events relating to the same
// contact point can be identified by the ID of the sequence.
type Touch struct {
    ObjectID
    listener TouchListener
}

func (this *Touch) AddListener(listener TouchListener) {
    this.listener = listener
}

func (this *Touch) Release() error {
    return nil
}




const OutputSubpixelUnknown = 0 // unknown geometry
const OutputSubpixelNone = 1 // no geometry
const OutputSubpixelHorizontalRgb = 2 // horizontal RGB
const OutputSubpixelHorizontalBgr = 3 // horizontal BGR
const OutputSubpixelVerticalRgb = 4 // vertical RGB
const OutputSubpixelVerticalBgr = 5 // vertical BGR


const OutputTransformNormal = 0 // no transform
const OutputTransform90 = 1 // 90 degrees counter-clockwise
const OutputTransform180 = 2 // 180 degrees counter-clockwise
const OutputTransform270 = 3 // 270 degrees counter-clockwise
const OutputTransformFlipped = 4 // 180 degree flip around a vertical axis
const OutputTransformFlipped90 = 5 // flip and rotate 90 degrees counter-clockwise
const OutputTransformFlipped180 = 6 // flip and rotate 180 degrees counter-clockwise
const OutputTransformFlipped270 = 7 // flip and rotate 270 degrees counter-clockwise


const OutputModeCurrent = 0x1 // indicates this is the current mode
const OutputModePreferred = 0x2 // indicates this is the preferred mode

type OutputListener interface {
    Geometry(x int32, y int32, physicalWidth int32, physicalHeight int32, subpixel int32, make string, model string, transform int32)
    Mode(flags uint32, width int32, height int32, refresh int32)
    Done()
    Scale(factor int32)
}

// An output describes part of the compositor geometry.  The
// compositor works in the 'compositor coordinate system' and an
// output corresponds to a rectangular area in that space that is
// actually visible.  This typically corresponds to a monitor that
// displays part of the compositor space.  This object is published
// as global during start up, or when a monitor is hotplugged.
type Output struct {
    ObjectID
    listener OutputListener
}

func (this *Output) AddListener(listener OutputListener) {
    this.listener = listener
}

// Using this request a client can tell the server that it is not going to
// use the output object anymore.
func (this *Output) Release() error {
    return nil
}



type RegionListener interface {
}

// A region object describes an area.
// 
// Region objects are used to describe the opaque and input
// regions of a surface.
type Region struct {
    ObjectID
    listener RegionListener
}

func (this *Region) AddListener(listener RegionListener) {
    this.listener = listener
}

// Destroy the region.  This will invalidate the object ID.
func (this *Region) Destroy() error {
    return nil
}

// Add the specified rectangle to the region.
func (this *Region) Add(x int32, y int32, width int32, height int32) error {
    return nil
}

// Subtract the specified rectangle from the region.
func (this *Region) Subtract(x int32, y int32, width int32, height int32) error {
    return nil
}




const SubcompositorErrorBadSurface = 0 // the to-be sub-surface is invalid

type SubcompositorListener interface {
}

// The global interface exposing sub-surface compositing capabilities.
// A wl_surface, that has sub-surfaces associated, is called the
// parent surface. Sub-surfaces can be arbitrarily nested and create
// a tree of sub-surfaces.
// 
// The root surface in a tree of sub-surfaces is the main
// surface. The main surface cannot be a sub-surface, because
// sub-surfaces must always have a parent.
// 
// A main surface with its sub-surfaces forms a (compound) window.
// For window management purposes, this set of wl_surface objects is
// to be considered as a single window, and it should also behave as
// such.
// 
// The aim of sub-surfaces is to offload some of the compositing work
// within a window from clients to the compositor. A prime example is
// a video player with decorations and video in separate wl_surface
// objects. This should allow the compositor to pass YUV video buffer
// processing to dedicated overlay hardware when possible.
type Subcompositor struct {
    ObjectID
    listener SubcompositorListener
}

func (this *Subcompositor) AddListener(listener SubcompositorListener) {
    this.listener = listener
}

// Informs the server that the client will not be using this
// protocol object anymore. This does not affect any other
// objects, wl_subsurface objects included.
func (this *Subcompositor) Destroy() error {
    return nil
}

// Create a sub-surface interface for the given surface, and
// associate it with the given parent surface. This turns a
// plain wl_surface into a sub-surface.
// 
// The to-be sub-surface must not already have another role, and it
// must not have an existing wl_subsurface object. Otherwise a protocol
// error is raised.
func (this *Subcompositor) GetSubsurface(surface uint32, parent uint32) (*Subsurface, error) {
    return nil, nil
}




const SubsurfaceErrorBadSurface = 0 // wl_surface is not a sibling or the parent

type SubsurfaceListener interface {
}

// An additional interface to a wl_surface object, which has been
// made a sub-surface. A sub-surface has one parent surface. A
// sub-surface's size and position are not limited to that of the parent.
// Particularly, a sub-surface is not automatically clipped to its
// parent's area.
// 
// A sub-surface becomes mapped, when a non-NULL wl_buffer is applied
// and the parent surface is mapped. The order of which one happens
// first is irrelevant. A sub-surface is hidden if the parent becomes
// hidden, or if a NULL wl_buffer is applied. These rules apply
// recursively through the tree of surfaces.
// 
// The behaviour of a wl_surface.commit request on a sub-surface
// depends on the sub-surface's mode. The possible modes are
// synchronized and desynchronized, see methods
// wl_subsurface.set_sync and wl_subsurface.set_desync. Synchronized
// mode caches the wl_surface state to be applied when the parent's
// state gets applied, and desynchronized mode applies the pending
// wl_surface state directly. A sub-surface is initially in the
// synchronized mode.
// 
// Sub-surfaces have also other kind of state, which is managed by
// wl_subsurface requests, as opposed to wl_surface requests. This
// state includes the sub-surface position relative to the parent
// surface (wl_subsurface.set_position), and the stacking order of
// the parent and its sub-surfaces (wl_subsurface.place_above and
// .place_below). This state is applied when the parent surface's
// wl_surface state is applied, regardless of the sub-surface's mode.
// As the exception, set_sync and set_desync are effective immediately.
// 
// The main surface can be thought to be always in desynchronized mode,
// since it does not have a parent in the sub-surfaces sense.
// 
// Even if a sub-surface is in desynchronized mode, it will behave as
// in synchronized mode, if its parent surface behaves as in
// synchronized mode. This rule is applied recursively throughout the
// tree of surfaces. This means, that one can set a sub-surface into
// synchronized mode, and then assume that all its child and grand-child
// sub-surfaces are synchronized, too, without explicitly setting them.
// 
// If the wl_surface associated with the wl_subsurface is destroyed, the
// wl_subsurface object becomes inert. Note, that destroying either object
// takes effect immediately. If you need to synchronize the removal
// of a sub-surface to the parent surface update, unmap the sub-surface
// first by attaching a NULL wl_buffer, update parent, and then destroy
// the sub-surface.
// 
// If the parent wl_surface object is destroyed, the sub-surface is
// unmapped.
type Subsurface struct {
    ObjectID
    listener SubsurfaceListener
}

func (this *Subsurface) AddListener(listener SubsurfaceListener) {
    this.listener = listener
}

// The sub-surface interface is removed from the wl_surface object
// that was turned into a sub-surface with a
// wl_subcompositor.get_subsurface request. The wl_surface's association
// to the parent is deleted, and the wl_surface loses its role as
// a sub-surface. The wl_surface is unmapped.
func (this *Subsurface) Destroy() error {
    return nil
}

// This schedules a sub-surface position change.
// The sub-surface will be moved so that its origin (top left
// corner pixel) will be at the location x, y of the parent surface
// coordinate system. The coordinates are not restricted to the parent
// surface area. Negative values are allowed.
// 
// The scheduled coordinates will take effect whenever the state of the
// parent surface is applied. When this happens depends on whether the
// parent surface is in synchronized mode or not. See
// wl_subsurface.set_sync and wl_subsurface.set_desync for details.
// 
// If more than one set_position request is invoked by the client before
// the commit of the parent surface, the position of a new request always
// replaces the scheduled position from any previous request.
// 
// The initial position is 0, 0.
func (this *Subsurface) SetPosition(x int32, y int32) error {
    return nil
}

// This sub-surface is taken from the stack, and put back just
// above the reference surface, changing the z-order of the sub-surfaces.
// The reference surface must be one of the sibling surfaces, or the
// parent surface. Using any other surface, including this sub-surface,
// will cause a protocol error.
// 
// The z-order is double-buffered. Requests are handled in order and
// applied immediately to a pending state. The final pending state is
// copied to the active state the next time the state of the parent
// surface is applied. When this happens depends on whether the parent
// surface is in synchronized mode or not. See wl_subsurface.set_sync and
// wl_subsurface.set_desync for details.
// 
// A new sub-surface is initially added as the top-most in the stack
// of its siblings and parent.
func (this *Subsurface) PlaceAbove(sibling uint32) error {
    return nil
}

// The sub-surface is placed just below the reference surface.
// See wl_subsurface.place_above.
func (this *Subsurface) PlaceBelow(sibling uint32) error {
    return nil
}

// Change the commit behaviour of the sub-surface to synchronized
// mode, also described as the parent dependent mode.
// 
// In synchronized mode, wl_surface.commit on a sub-surface will
// accumulate the committed state in a cache, but the state will
// not be applied and hence will not change the compositor output.
// The cached state is applied to the sub-surface immediately after
// the parent surface's state is applied. This ensures atomic
// updates of the parent and all its synchronized sub-surfaces.
// Applying the cached state will invalidate the cache, so further
// parent surface commits do not (re-)apply old state.
// 
// See wl_subsurface for the recursive effect of this mode.
func (this *Subsurface) SetSync() error {
    return nil
}

// Change the commit behaviour of the sub-surface to desynchronized
// mode, also described as independent or freely running mode.
// 
// In desynchronized mode, wl_surface.commit on a sub-surface will
// apply the pending state directly, without caching, as happens
// normally with a wl_surface. Calling wl_surface.commit on the
// parent surface has no effect on the sub-surface's wl_surface
// state. This mode allows a sub-surface to be updated on its own.
// 
// If cached state exists when wl_surface.commit is called in
// desynchronized mode, the pending state is added to the cached
// state, and applied as a whole. This invalidates the cache.
// 
// Note: even if a sub-surface is set to desynchronized, a parent
// sub-surface may override it to behave as synchronized. For details,
// see wl_subsurface.
// 
// If a surface's parent surface behaves as desynchronized, then
// the cached state is applied on set_desync.
func (this *Subsurface) SetDesync() error {
    return nil
}

