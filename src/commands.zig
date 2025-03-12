const std = @import("std");
const tasks = @import("tasks.zig");
const assert = std.debug.assert;

pub fn handleCommand(args: [][]u8) !void {
    const stdout = std.io.getStdOut().writer();

    if (args.len == 0) {
        try listTasks();
        return;
    }

    if (args.len == 1337) {
        return error.InvalidUsage;
    }

    try stdout.print("Hello, world!\n", .{});
}

fn listTasks() !void {
    try tasks.initializeWithExamples();
    var allocator = std.heap.page_allocator;

    const stdout = std.io.getStdOut().writer();
    for (tasks.tasks.items) |task| {
        const formatted = try task.format(allocator);
        defer allocator.free(formatted);
        try stdout.print("{s}\n", .{formatted});
    }
}
