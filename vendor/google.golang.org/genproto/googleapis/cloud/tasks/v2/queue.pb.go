// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2/queue.proto

package tasks

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// State of the queue.
type Queue_State int32

const (
	// Unspecified state.
	Queue_STATE_UNSPECIFIED Queue_State = 0
	// The queue is running. Tasks can be dispatched.
	//
	// If the queue was created using Cloud Tasks and the queue has
	// had no activity (method calls or task dispatches) for 30 days,
	// the queue may take a few minutes to re-activate. Some method
	// calls may return [NOT_FOUND][google.rpc.Code.NOT_FOUND] and
	// tasks may not be dispatched for a few minutes until the queue
	// has been re-activated.
	Queue_RUNNING Queue_State = 1
	// Tasks are paused by the user. If the queue is paused then Cloud
	// Tasks will stop delivering tasks from it, but more tasks can
	// still be added to it by the user.
	Queue_PAUSED Queue_State = 2
	// The queue is disabled.
	//
	// A queue becomes `DISABLED` when
	// [queue.yaml](https://cloud.google.com/appengine/docs/python/config/queueref)
	// or
	// [queue.xml](https://cloud.google.com/appengine/docs/standard/java/config/queueref)
	// is uploaded which does not contain the queue. You cannot directly disable
	// a queue.
	//
	// When a queue is disabled, tasks can still be added to a queue
	// but the tasks are not dispatched.
	//
	// To permanently delete this queue and all of its tasks, call
	// [DeleteQueue][google.cloud.tasks.v2.CloudTasks.DeleteQueue].
	Queue_DISABLED Queue_State = 3
)

var Queue_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "RUNNING",
	2: "PAUSED",
	3: "DISABLED",
}

var Queue_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"RUNNING":           1,
	"PAUSED":            2,
	"DISABLED":          3,
}

func (x Queue_State) String() string {
	return proto.EnumName(Queue_State_name, int32(x))
}

func (Queue_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{0, 0}
}

// A queue is a container of related tasks. Queues are configured to manage
// how those tasks are dispatched. Configurable properties include rate limits,
// retry options, queue types, and others.
type Queue struct {
	// Caller-specified and required in [CreateQueue][google.cloud.tasks.v2.CloudTasks.CreateQueue],
	// after which it becomes output only.
	//
	// The queue name.
	//
	// The queue name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	//    projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the queue's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Overrides for
	// [task-level app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	// These settings apply only to
	// [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest] in this queue.
	// [Http tasks][google.cloud.tasks.v2.HttpRequest] are not affected.
	//
	// If set, `app_engine_routing_override` is used for all
	// [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest] in the queue, no matter what the
	// setting is for the
	// [task-level app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	AppEngineRoutingOverride *AppEngineRouting `protobuf:"bytes,2,opt,name=app_engine_routing_override,json=appEngineRoutingOverride,proto3" json:"app_engine_routing_override,omitempty"`
	// Rate limits for task dispatches.
	//
	// [rate_limits][google.cloud.tasks.v2.Queue.rate_limits] and [retry_config][google.cloud.tasks.v2.Queue.retry_config] are
	// related because they both control task attempts. However they control task
	// attempts in different ways:
	//
	// * [rate_limits][google.cloud.tasks.v2.Queue.rate_limits] controls the total rate of
	//   dispatches from a queue (i.e. all traffic dispatched from the
	//   queue, regardless of whether the dispatch is from a first
	//   attempt or a retry).
	// * [retry_config][google.cloud.tasks.v2.Queue.retry_config] controls what happens to
	//   particular a task after its first attempt fails. That is,
	//   [retry_config][google.cloud.tasks.v2.Queue.retry_config] controls task retries (the
	//   second attempt, third attempt, etc).
	//
	// The queue's actual dispatch rate is the result of:
	//
	// * Number of tasks in the queue
	// * User-specified throttling: [rate_limits][google.cloud.tasks.v2.Queue.rate_limits],
	//   [retry_config][google.cloud.tasks.v2.Queue.retry_config], and the
	//   [queue's state][google.cloud.tasks.v2.Queue.state].
	// * System throttling due to `429` (Too Many Requests) or `503` (Service
	//   Unavailable) responses from the worker, high error rates, or to smooth
	//   sudden large traffic spikes.
	RateLimits *RateLimits `protobuf:"bytes,3,opt,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Settings that determine the retry behavior.
	//
	// * For tasks created using Cloud Tasks: the queue-level retry settings
	//   apply to all tasks in the queue that were created using Cloud Tasks.
	//   Retry settings cannot be set on individual tasks.
	// * For tasks created using the App Engine SDK: the queue-level retry
	//   settings apply to all tasks in the queue which do not have retry settings
	//   explicitly set on the task and were created by the App Engine SDK. See
	//   [App Engine
	//   documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
	RetryConfig *RetryConfig `protobuf:"bytes,4,opt,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty"`
	// Output only. The state of the queue.
	//
	// `state` can only be changed by called
	// [PauseQueue][google.cloud.tasks.v2.CloudTasks.PauseQueue],
	// [ResumeQueue][google.cloud.tasks.v2.CloudTasks.ResumeQueue], or uploading
	// [queue.yaml/xml](https://cloud.google.com/appengine/docs/python/config/queueref).
	// [UpdateQueue][google.cloud.tasks.v2.CloudTasks.UpdateQueue] cannot be used to change `state`.
	State Queue_State `protobuf:"varint,5,opt,name=state,proto3,enum=google.cloud.tasks.v2.Queue_State" json:"state,omitempty"`
	// Output only. The last time this queue was purged.
	//
	// All tasks that were [created][google.cloud.tasks.v2.Task.create_time] before this time
	// were purged.
	//
	// A queue can be purged using [PurgeQueue][google.cloud.tasks.v2.CloudTasks.PurgeQueue], the
	// [App Engine Task Queue SDK, or the Cloud
	// Console](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/deleting-tasks-and-queues#purging_all_tasks_from_a_queue).
	//
	// Purge time will be truncated to the nearest microsecond. Purge
	// time will be unset if the queue has never been purged.
	PurgeTime            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=purge_time,json=purgeTime,proto3" json:"purge_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Queue) Reset()         { *m = Queue{} }
func (m *Queue) String() string { return proto.CompactTextString(m) }
func (*Queue) ProtoMessage()    {}
func (*Queue) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{0}
}

func (m *Queue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Queue.Unmarshal(m, b)
}
func (m *Queue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Queue.Marshal(b, m, deterministic)
}
func (m *Queue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Queue.Merge(m, src)
}
func (m *Queue) XXX_Size() int {
	return xxx_messageInfo_Queue.Size(m)
}
func (m *Queue) XXX_DiscardUnknown() {
	xxx_messageInfo_Queue.DiscardUnknown(m)
}

var xxx_messageInfo_Queue proto.InternalMessageInfo

func (m *Queue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Queue) GetAppEngineRoutingOverride() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRoutingOverride
	}
	return nil
}

func (m *Queue) GetRateLimits() *RateLimits {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *Queue) GetRetryConfig() *RetryConfig {
	if m != nil {
		return m.RetryConfig
	}
	return nil
}

func (m *Queue) GetState() Queue_State {
	if m != nil {
		return m.State
	}
	return Queue_STATE_UNSPECIFIED
}

func (m *Queue) GetPurgeTime() *timestamp.Timestamp {
	if m != nil {
		return m.PurgeTime
	}
	return nil
}

// Rate limits.
//
// This message determines the maximum rate that tasks can be dispatched by a
// queue, regardless of whether the dispatch is a first task attempt or a retry.
//
// Note: The debugging command, [RunTask][google.cloud.tasks.v2.CloudTasks.RunTask], will run a task
// even if the queue has reached its [RateLimits][google.cloud.tasks.v2.RateLimits].
type RateLimits struct {
	// The maximum rate at which tasks are dispatched from this queue.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// * The maximum allowed value is 500.
	//
	//
	// This field has the same meaning as
	// [rate in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#rate).
	MaxDispatchesPerSecond float64 `protobuf:"fixed64,1,opt,name=max_dispatches_per_second,json=maxDispatchesPerSecond,proto3" json:"max_dispatches_per_second,omitempty"`
	// Output only. The max burst size.
	//
	// Max burst size limits how fast tasks in queue are processed when
	// many tasks are in the queue and the rate is high. This field
	// allows the queue to have a high rate so processing starts shortly
	// after a task is enqueued, but still limits resource usage when
	// many tasks are enqueued in a short period of time.
	//
	// The [token bucket](https://wikipedia.org/wiki/Token_Bucket)
	// algorithm is used to control the rate of task dispatches. Each
	// queue has a token bucket that holds tokens, up to the maximum
	// specified by `max_burst_size`. Each time a task is dispatched, a
	// token is removed from the bucket. Tasks will be dispatched until
	// the queue's bucket runs out of tokens. The bucket will be
	// continuously refilled with new tokens based on
	// [max_dispatches_per_second][google.cloud.tasks.v2.RateLimits.max_dispatches_per_second].
	//
	// Cloud Tasks will pick the value of `max_burst_size` based on the
	// value of
	// [max_dispatches_per_second][google.cloud.tasks.v2.RateLimits.max_dispatches_per_second].
	//
	// For queues that were created or updated using
	// `queue.yaml/xml`, `max_burst_size` is equal to
	// [bucket_size](https://cloud.google.com/appengine/docs/standard/python/config/queueref#bucket_size).
	// Since `max_burst_size` is output only, if
	// [UpdateQueue][google.cloud.tasks.v2.CloudTasks.UpdateQueue] is called on a queue
	// created by `queue.yaml/xml`, `max_burst_size` will be reset based
	// on the value of
	// [max_dispatches_per_second][google.cloud.tasks.v2.RateLimits.max_dispatches_per_second],
	// regardless of whether
	// [max_dispatches_per_second][google.cloud.tasks.v2.RateLimits.max_dispatches_per_second]
	// is updated.
	//
	MaxBurstSize int32 `protobuf:"varint,2,opt,name=max_burst_size,json=maxBurstSize,proto3" json:"max_burst_size,omitempty"`
	// The maximum number of concurrent tasks that Cloud Tasks allows
	// to be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// The maximum allowed value is 5,000.
	//
	//
	// This field has the same meaning as
	// [max_concurrent_requests in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#max_concurrent_requests).
	MaxConcurrentDispatches int32    `protobuf:"varint,3,opt,name=max_concurrent_dispatches,json=maxConcurrentDispatches,proto3" json:"max_concurrent_dispatches,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *RateLimits) Reset()         { *m = RateLimits{} }
func (m *RateLimits) String() string { return proto.CompactTextString(m) }
func (*RateLimits) ProtoMessage()    {}
func (*RateLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{1}
}

func (m *RateLimits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimits.Unmarshal(m, b)
}
func (m *RateLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimits.Marshal(b, m, deterministic)
}
func (m *RateLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimits.Merge(m, src)
}
func (m *RateLimits) XXX_Size() int {
	return xxx_messageInfo_RateLimits.Size(m)
}
func (m *RateLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimits.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimits proto.InternalMessageInfo

func (m *RateLimits) GetMaxDispatchesPerSecond() float64 {
	if m != nil {
		return m.MaxDispatchesPerSecond
	}
	return 0
}

func (m *RateLimits) GetMaxBurstSize() int32 {
	if m != nil {
		return m.MaxBurstSize
	}
	return 0
}

func (m *RateLimits) GetMaxConcurrentDispatches() int32 {
	if m != nil {
		return m.MaxConcurrentDispatches
	}
	return 0
}

// Retry config.
//
// These settings determine when a failed task attempt is retried.
type RetryConfig struct {
	// Number of attempts per task.
	//
	// Cloud Tasks will attempt the task `max_attempts` times (that is, if the
	// first attempt fails, then there will be `max_attempts - 1` retries). Must
	// be >= -1.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// -1 indicates unlimited attempts.
	//
	// This field has the same meaning as
	// [task_retry_limit in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxAttempts int32 `protobuf:"varint,1,opt,name=max_attempts,json=maxAttempts,proto3" json:"max_attempts,omitempty"`
	// If positive, `max_retry_duration` specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once `max_retry_duration` time has passed *and* the
	// task has been attempted [max_attempts][google.cloud.tasks.v2.RetryConfig.max_attempts]
	// times, no further attempts will be made and the task will be
	// deleted.
	//
	// If zero, then the task age is unlimited.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// `max_retry_duration` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [task_age_limit in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxRetryDuration *duration.Duration `protobuf:"bytes,2,opt,name=max_retry_duration,json=maxRetryDuration,proto3" json:"max_retry_duration,omitempty"`
	// A task will be [scheduled][google.cloud.tasks.v2.Task.schedule_time] for retry between
	// [min_backoff][google.cloud.tasks.v2.RetryConfig.min_backoff] and
	// [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] duration after it fails,
	// if the queue's [RetryConfig][google.cloud.tasks.v2.RetryConfig] specifies that the task should be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// `min_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [min_backoff_seconds in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MinBackoff *duration.Duration `protobuf:"bytes,3,opt,name=min_backoff,json=minBackoff,proto3" json:"min_backoff,omitempty"`
	// A task will be [scheduled][google.cloud.tasks.v2.Task.schedule_time] for retry between
	// [min_backoff][google.cloud.tasks.v2.RetryConfig.min_backoff] and
	// [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] duration after it fails,
	// if the queue's [RetryConfig][google.cloud.tasks.v2.RetryConfig] specifies that the task should be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// `max_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [max_backoff_seconds in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxBackoff *duration.Duration `protobuf:"bytes,4,opt,name=max_backoff,json=maxBackoff,proto3" json:"max_backoff,omitempty"`
	// The time between retries will double `max_doublings` times.
	//
	// A task's retry interval starts at
	// [min_backoff][google.cloud.tasks.v2.RetryConfig.min_backoff], then doubles
	// `max_doublings` times, then increases linearly, and finally
	// retries retries at intervals of
	// [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] up to
	// [max_attempts][google.cloud.tasks.v2.RetryConfig.max_attempts] times.
	//
	// For example, if [min_backoff][google.cloud.tasks.v2.RetryConfig.min_backoff] is 10s,
	// [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] is 300s, and
	// `max_doublings` is 3, then the a task will first be retried in
	// 10s. The retry interval will double three times, and then
	// increase linearly by 2^3 * 10s.  Finally, the task will retry at
	// intervals of [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] until the
	// task has been attempted [max_attempts][google.cloud.tasks.v2.RetryConfig.max_attempts]
	// times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s,
	// 240s, 300s, 300s, ....
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// This field has the same meaning as
	// [max_doublings in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxDoublings         int32    `protobuf:"varint,5,opt,name=max_doublings,json=maxDoublings,proto3" json:"max_doublings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryConfig) Reset()         { *m = RetryConfig{} }
func (m *RetryConfig) String() string { return proto.CompactTextString(m) }
func (*RetryConfig) ProtoMessage()    {}
func (*RetryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{2}
}

func (m *RetryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryConfig.Unmarshal(m, b)
}
func (m *RetryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryConfig.Marshal(b, m, deterministic)
}
func (m *RetryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryConfig.Merge(m, src)
}
func (m *RetryConfig) XXX_Size() int {
	return xxx_messageInfo_RetryConfig.Size(m)
}
func (m *RetryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RetryConfig proto.InternalMessageInfo

func (m *RetryConfig) GetMaxAttempts() int32 {
	if m != nil {
		return m.MaxAttempts
	}
	return 0
}

func (m *RetryConfig) GetMaxRetryDuration() *duration.Duration {
	if m != nil {
		return m.MaxRetryDuration
	}
	return nil
}

func (m *RetryConfig) GetMinBackoff() *duration.Duration {
	if m != nil {
		return m.MinBackoff
	}
	return nil
}

func (m *RetryConfig) GetMaxBackoff() *duration.Duration {
	if m != nil {
		return m.MaxBackoff
	}
	return nil
}

func (m *RetryConfig) GetMaxDoublings() int32 {
	if m != nil {
		return m.MaxDoublings
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.cloud.tasks.v2.Queue_State", Queue_State_name, Queue_State_value)
	proto.RegisterType((*Queue)(nil), "google.cloud.tasks.v2.Queue")
	proto.RegisterType((*RateLimits)(nil), "google.cloud.tasks.v2.RateLimits")
	proto.RegisterType((*RetryConfig)(nil), "google.cloud.tasks.v2.RetryConfig")
}

func init() { proto.RegisterFile("google/cloud/tasks/v2/queue.proto", fileDescriptor_a4a1833e2495b95c) }

var fileDescriptor_a4a1833e2495b95c = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x61, 0x4f, 0x13, 0x4d,
	0x10, 0x7e, 0x0f, 0x5a, 0x5e, 0x99, 0x22, 0xa9, 0x9b, 0xa0, 0x2d, 0x1a, 0x81, 0x6a, 0x22, 0x9f,
	0xee, 0x12, 0x4c, 0x8c, 0xd4, 0x4f, 0x2d, 0x3d, 0x49, 0x13, 0x52, 0xeb, 0x15, 0x3e, 0x68, 0x4c,
	0x2e, 0xdb, 0xeb, 0xf6, 0x5c, 0xe9, 0xee, 0x9e, 0xbb, 0x7b, 0xa4, 0x42, 0xf8, 0x1b, 0xfe, 0x06,
	0xfd, 0x6b, 0xfe, 0x0a, 0x73, 0x73, 0x77, 0x94, 0x20, 0xe8, 0xa7, 0xce, 0xce, 0x3c, 0xcf, 0x33,
	0xbb, 0x33, 0x4f, 0x0f, 0x76, 0x62, 0xa5, 0xe2, 0x19, 0xf3, 0xa2, 0x99, 0x4a, 0x27, 0x9e, 0xa5,
	0xe6, 0xd4, 0x78, 0x67, 0x7b, 0xde, 0xd7, 0x94, 0xa5, 0xcc, 0x4d, 0xb4, 0xb2, 0x8a, 0x6c, 0xe4,
	0x10, 0x17, 0x21, 0x2e, 0x42, 0xdc, 0xb3, 0xbd, 0xcd, 0x66, 0xc1, 0xa4, 0x09, 0xf7, 0x34, 0x33,
	0x2a, 0xd5, 0x51, 0xc1, 0xd8, 0x6c, 0xdd, 0x2e, 0x6a, 0xa9, 0x8e, 0x99, 0x2d, 0x30, 0x4f, 0x0b,
	0x0c, 0x9e, 0xc6, 0xe9, 0xd4, 0x9b, 0xa4, 0x9a, 0x5a, 0xae, 0x64, 0x51, 0xdf, 0xba, 0x59, 0xb7,
	0x5c, 0x30, 0x63, 0xa9, 0x48, 0x0a, 0xc0, 0x93, 0x6b, 0xfd, 0xa9, 0x94, 0xca, 0x22, 0xdb, 0xe4,
	0xd5, 0xd6, 0x8f, 0x0a, 0x54, 0xdf, 0x67, 0x8f, 0x20, 0x04, 0x2a, 0x92, 0x0a, 0xd6, 0x70, 0xb6,
	0x9d, 0xdd, 0xd5, 0x00, 0x63, 0x32, 0x85, 0xc7, 0x34, 0x49, 0x42, 0x26, 0x63, 0x2e, 0x59, 0xa8,
	0x55, 0x6a, 0xb9, 0x8c, 0x43, 0x75, 0xc6, 0xb4, 0xe6, 0x13, 0xd6, 0x58, 0xda, 0x76, 0x76, 0x6b,
	0x7b, 0x2f, 0xdc, 0x5b, 0x1f, 0xee, 0x76, 0x92, 0xc4, 0x47, 0x62, 0x90, 0xf3, 0x82, 0x06, 0xbd,
	0x91, 0x79, 0x57, 0x08, 0x91, 0x2e, 0xd4, 0x34, 0xb5, 0x2c, 0x9c, 0x71, 0xc1, 0xad, 0x69, 0x2c,
	0xa3, 0xee, 0xce, 0x1d, 0xba, 0x01, 0xb5, 0xec, 0x08, 0x81, 0x01, 0xe8, 0xab, 0x98, 0xf8, 0xb0,
	0xa6, 0x99, 0xd5, 0xdf, 0xc2, 0x48, 0xc9, 0x29, 0x8f, 0x1b, 0x15, 0x14, 0x69, 0xdd, 0x25, 0x92,
	0x41, 0x0f, 0x10, 0x19, 0xd4, 0xf4, 0xe2, 0x40, 0x5e, 0x43, 0xd5, 0x58, 0x6a, 0x59, 0xa3, 0xba,
	0xed, 0xec, 0xae, 0xdf, 0xc9, 0xc7, 0x99, 0xb9, 0xa3, 0x0c, 0x19, 0xe4, 0x04, 0xb2, 0x0f, 0x90,
	0xa4, 0x3a, 0x66, 0x61, 0xb6, 0x81, 0xc6, 0x0a, 0xb6, 0xdf, 0x2c, 0xe9, 0xe5, 0x7a, 0xdc, 0xe3,
	0x72, 0x3d, 0xc1, 0x2a, 0xa2, 0xb3, 0x73, 0xcb, 0x87, 0x2a, 0x4a, 0x91, 0x0d, 0x78, 0x30, 0x3a,
	0xee, 0x1c, 0xfb, 0xe1, 0xc9, 0x60, 0x34, 0xf4, 0x0f, 0xfa, 0x6f, 0xfb, 0x7e, 0xaf, 0xfe, 0x1f,
	0xa9, 0xc1, 0xff, 0xc1, 0xc9, 0x60, 0xd0, 0x1f, 0x1c, 0xd6, 0x1d, 0x02, 0xb0, 0x32, 0xec, 0x9c,
	0x8c, 0xfc, 0x5e, 0x7d, 0x89, 0xac, 0xc1, 0xbd, 0x5e, 0x7f, 0xd4, 0xe9, 0x1e, 0xf9, 0xbd, 0xfa,
	0x72, 0xfb, 0xd3, 0xaf, 0xce, 0x07, 0xd8, 0xc2, 0x9b, 0xe6, 0x17, 0xcd, 0x9b, 0xd3, 0x84, 0x1b,
	0x37, 0x52, 0xc2, 0xcb, 0x17, 0xfd, 0x2a, 0xd1, 0xea, 0x0b, 0x8b, 0xac, 0xf1, 0x2e, 0x8a, 0xe8,
	0xd2, 0x9b, 0xa9, 0x28, 0xb7, 0x85, 0x77, 0x51, 0x86, 0x97, 0xb9, 0xb9, 0x8d, 0x77, 0x81, 0xbf,
	0x97, 0xad, 0x9f, 0x0e, 0xc0, 0x62, 0xf6, 0x64, 0x1f, 0x9a, 0x82, 0xce, 0xc3, 0x09, 0x37, 0x09,
	0xb5, 0xd1, 0x67, 0x66, 0xc2, 0x84, 0xe9, 0xd0, 0xb0, 0x48, 0xc9, 0x09, 0x9a, 0xc8, 0x09, 0x1e,
	0x0a, 0x3a, 0xef, 0x5d, 0xd5, 0x87, 0x4c, 0x8f, 0xb0, 0x4a, 0x9e, 0xc3, 0x7a, 0x46, 0x1d, 0xa7,
	0xda, 0xd8, 0xd0, 0xf0, 0xf3, 0xdc, 0x49, 0xd5, 0x60, 0x4d, 0xd0, 0x79, 0x37, 0x4b, 0x8e, 0xf8,
	0x39, 0x23, 0xed, 0xbc, 0x41, 0xa4, 0x64, 0x94, 0x6a, 0xcd, 0xa4, 0xbd, 0xd6, 0x0b, 0x2d, 0x52,
	0x0d, 0x1e, 0x09, 0x3a, 0x3f, 0xb8, 0xaa, 0x2f, 0x5a, 0xb5, 0xbe, 0x2f, 0x41, 0xed, 0xda, 0x8a,
	0xc9, 0x0e, 0x64, 0xda, 0x21, 0xb5, 0x96, 0x89, 0xc4, 0x1a, 0xbc, 0x5f, 0x35, 0xa8, 0x09, 0x3a,
	0xef, 0x14, 0x29, 0x72, 0x08, 0x24, 0x83, 0xe4, 0x1e, 0x2a, 0xff, 0x64, 0x85, 0xc5, 0x9b, 0x7f,
	0xac, 0xb1, 0x57, 0x00, 0x82, 0xba, 0xa0, 0x73, 0xec, 0x54, 0x66, 0x48, 0x1b, 0x6a, 0x82, 0xcb,
	0x70, 0x4c, 0xa3, 0x53, 0x35, 0x9d, 0x16, 0x66, 0xfe, 0x8b, 0x02, 0x08, 0x2e, 0xbb, 0x39, 0x18,
	0xb9, 0xd9, 0x64, 0x0a, 0x6e, 0xe5, 0xdf, 0x5c, 0x3a, 0x2f, 0xb9, 0xcf, 0xe0, 0x3e, 0x2e, 0x44,
	0xa5, 0xe3, 0x19, 0x97, 0xb1, 0x41, 0x07, 0xe7, 0x43, 0xed, 0x95, 0xb9, 0x2e, 0x83, 0x66, 0xa4,
	0xc4, 0xed, 0xa6, 0xee, 0x02, 0x1a, 0x64, 0x98, 0x75, 0x19, 0x3a, 0x1f, 0xdb, 0x05, 0x28, 0x56,
	0x33, 0x2a, 0x63, 0x57, 0xe9, 0xd8, 0x8b, 0x99, 0xc4, 0x3b, 0x78, 0x0b, 0x63, 0xdd, 0xf8, 0x78,
	0xbd, 0xc1, 0x60, 0xbc, 0x82, 0xb0, 0x97, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x1a, 0x7f,
	0xbb, 0x37, 0x05, 0x00, 0x00,
}