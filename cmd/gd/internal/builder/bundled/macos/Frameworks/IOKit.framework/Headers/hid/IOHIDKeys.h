// IOKit HID property keys for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_IOKIT_IOHIDKEYS_H
#define GD_IOKIT_IOHIDKEYS_H

#define kIOHIDDeviceKey "IOHIDDevice"
#define kIOHIDTransportKey "Transport"
#define kIOHIDVendorIDKey "VendorID"
#define kIOHIDVendorIDSourceKey "VendorIDSource"
#define kIOHIDProductIDKey "ProductID"
#define kIOHIDVersionNumberKey "VersionNumber"
#define kIOHIDManufacturerKey "Manufacturer"
#define kIOHIDProductKey "Product"
#define kIOHIDSerialNumberKey "SerialNumber"
#define kIOHIDCountryCodeKey "CountryCode"
#define kIOHIDLocationIDKey "LocationID"
#define kIOHIDDeviceUsageKey "DeviceUsage"
#define kIOHIDDeviceUsagePageKey "DeviceUsagePage"
#define kIOHIDDeviceUsagePairsKey "DeviceUsagePairs"
#define kIOHIDPrimaryUsageKey "PrimaryUsage"
#define kIOHIDPrimaryUsagePageKey "PrimaryUsagePage"
#define kIOHIDMaxInputReportSizeKey "MaxInputReportSize"
#define kIOHIDMaxOutputReportSizeKey "MaxOutputReportSize"
#define kIOHIDMaxFeatureReportSizeKey "MaxFeatureReportSize"
#define kIOHIDReportIntervalKey "ReportInterval"
#define kIOHIDReportDescriptorKey "ReportDescriptor"
#define kIOHIDUniqueIDKey "UniqueID"
#define kIOHIDPhysicalDeviceUniqueIDKey "PhysicalDeviceUniqueID"

#define kIOHIDTransportUSBValue "USB"
#define kIOHIDTransportBluetoothValue "Bluetooth"
#define kIOHIDTransportBluetoothLowEnergyValue "BluetoothLowEnergy"
#define kIOHIDTransportSPIValue "SPI"
#define kIOHIDTransportI2CValue "I2C"

#define kIOHIDElementKey "Elements"
#define kIOHIDElementCookieKey "ElementCookie"
#define kIOHIDElementTypeKey "Type"
#define kIOHIDElementUsageKey "Usage"
#define kIOHIDElementUsagePageKey "UsagePage"
#define kIOHIDElementMinKey "Min"
#define kIOHIDElementMaxKey "Max"

#endif
