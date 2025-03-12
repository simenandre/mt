const std = @import("std");

pub const Task = struct {
    description: []const u8,
    completed: bool,

    pub fn format(self: Task, allocator: std.mem.Allocator) ![]const u8 {
        const status = if (self.completed) "✔" else " ";
        const formatted = try std.fmt.allocPrint(allocator, "[{s}] {s}", .{
            status,
            self.description,
        });

        // Return the formatted string
        return formatted;
    }
};

pub var tasks: std.ArrayList(Task) = std.ArrayList(Task).init(std.heap.page_allocator);

pub fn initializeWithExamples() !void {
    try tasks.append(Task{ .description = "Do the laundry", .completed = false });
    try tasks.append(Task{ .description = "Buy groceries", .completed = true });
    try tasks.append(Task{ .description = "Read a book", .completed = false });
}
