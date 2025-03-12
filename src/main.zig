const std = @import("std");
const commands = @import("commands.zig");
const assert = std.debug.assert;

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    // Get the command-line arguments
    var args = try std.process.argsAlloc(allocator);
    defer std.process.argsFree(allocator, args);

    // Skip the first argument (program name)
    const cmd_args = args[1..];

    // Handle the commands
    commands.handleCommand(cmd_args) catch |err| {
        switch (err) {
            error.InvalidUsage => {
                std.debug.print("Invalid usage. Use `mt help` for more information.\n", .{});
            },
            else => {
                std.debug.print("An unexpected error occurred: {s}\n", .{@errorName(err)});
            },
        }
        return; // Exit after handling the error
    };
}
