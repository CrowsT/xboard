package resolver

// Schema for api
const Schema = `

schema {
    query: Query
    mutation: Mutation
}

type Query {

    # A post object.
    post(id: String!): Post!


    # A user profile object.
    profile(): User!


    # A slice of thread.
    threadSlice(tags: [String!], query: SliceQuery!): ThreadSlice!
    # A thread object.
    thread(id: String!): Thread!


    # The count of unread notifications.
    unreadNotiCount(): UnreadNotiCount!
    # Notifications for current user.
    notification(type: String!, query: SliceQuery!): NotiSlice!


    # Containing mainTags and tagTree.
    tags(): Tags!

}

type Mutation {

    # Publish a new post.
    pubPost(post: PostInput!): Post!


    # Register/Login via email address. An email containing login info will be sent to the provided email address.
    auth(email: String!): Boolean!
    # Set the Name of user.
    setName(name: String!): User!
    # Directly edit tags subscribed by user.
    syncTags(tags: [String]!): User!
    # Add tags subscribed by user.
    addSubbedTags(tags: [String!]!): User!
    # Delete tags subscribed by user.
    delSubbedTags(tags: [String!]!): User!


    # Publish a new thread.
    pubThread(thread: ThreadInput!): Thread!

}


# Input object describing a Post to be published.
input PostInput {
    threadID: String!
    anonymous: Boolean!
    content: String!
    # Set quoting PostIDs.
    quotes: [String!]
}

# Object describing a Post.
type Post {
    id: String!
    anonymous: Boolean!
    author: String!
    content: String!
    createTime: Time!
    quotes: [Post!]
    quoteCount: Int!
}

# PostSlice object is for selecting specific 'slice' of Post objects to return. Affects returned SliceInfo.
type PostSlice {
    posts: [Post]!
    sliceInfo: SliceInfo!
}


type User {
    email: String!
    # The Name of user. Required when not posting anonymously.
    name: String
    # Tags saved by user.
    tags: [String!]
}


# Construct a new thread.
input ThreadInput {
    # Toggle anonymousness. If true, a new ID will be generated in each thread.
    anonymous: Boolean!
    content: String!
    # Required. Only one mainTag is allowed.
    mainTag: String!
    # Optional, maximum of 4.
    subTags: [String!]
    # Optional. If not set, the title will be '无题'.
    title: String
}

type Thread {
    # UUID with 8 chars in length, and will increase to 9 after 30 years.
    id: String!
    # Thread was published anonymously or not.
    anonymous: Boolean!
    # Same format as id if anonymous, name of User otherwise.
    author: String!
    content: String!
    createTime: Time!
    # Only one mainTag is allowed.
    mainTag: String!
    # Optional, maximum of 4.
    subTags: [String!]
    # Default to '无题'.
    title: String
    replies(query: SliceQuery!): PostSlice!
    replyCount: Int!
}

type ThreadSlice {
    threads: [Thread]!
    sliceInfo: SliceInfo!
}


# Count of different types of unread notifications.
type UnreadNotiCount {
    # Announcement messages from server.
    system: Int!
    # Threads that are replied.
    replied: Int!
    # Posts that are quoted.
    quoted: Int!
}

# NotiSlice object is for selecting specific 'slice' of an object to return. Affects returned SliceInfo.
type NotiSlice {
    # Announcement messages from server.
    system: [SystemNoti!]
    # Threads that are replied.
    replied: [RepliedNoti!]
    # Posts that are quoted.
    quoted: [Quoted!]
    # SliceInfo objects are generated by the server. Can be used in consecutive queries.
    sliceInfo: SliceInfo!
}

# Object describing a system notification.
type SystemNoti {
    # ID contains different types of format.
    id: String!
    # Type of Notification. "system", "replied" or "quoted".
    type: String!
    # Time when a notify event triggers. E.g. The time when a system event is announced from the server.
    eventTime: Time!
    # The nofitication is read or not.
    hasRead: Boolean!
    # Notification title.
    title: String!
    # Notification content.
    content: String!
}

# Object describing a replied notification.
type RepliedNoti {
    # ID contains different types of format.
    id: String!
    # Type of Notification. "system", "replied" or "quoted".
    type: String!
    # Time when a notify event triggers. E.g. The time when a thread is replied.
    eventTime: Time!
    # The nofitication is read or not.
    hasRead: Boolean!
    # The thread object that is replied.
    thread: Thread!
    # Array of users that replied. Same as the corresponding author field in the object "Post".
    repliers: [String!]!
}

# Object describing a quoted notification.
type Quoted {
    # ID contains different types of format.
    id: String!
    # Type of Notification. "system", "replied" or "quoted".
    type: String!
    # Time when a notify event triggers. E.g. The time when a post is quoted.
    eventTime: Time!
    hasRead: Boolean!
    # The thread object that is quoted in.
    thread: Thread!
    # The post object that is quoted.
    post: Post!
    # Array of users that quoted the post. Same as the corresponding author field in the object "Post".
    quoters: [String!]!
}


scalar Time

# SliceInfo objects are generated by the server. Can be used in consecutive queries.
type SliceInfo {
    firstCursor: String!
    lastCursor: String!
}

# SliceQuery object is for selecting specific 'slice' of an object to return. Affects returned SliceInfo.
input SliceQuery {
    # Either this field or 'after' is required.
    # An empty string means slice from the beginning.
    before: String
    # Either this field or 'before' is required.
    # An empty string means slice to the end.
    after: String
    # Set the amount of returned items.
    limit: Int!
}


type Tags {
    # Main tags are predefined manually.
    mainTags: [String!]!
    # Recommended tags are picked manually.
    recommended: [String!]!
    tree(query: String): [TagTreeNode!]
}

type TagTreeNode {
    mainTag: String!
    subTags: [String!]
}

`