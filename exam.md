### 1
- a

```sql
SELECT c.id, SUM((c.commission_rate*o.amount)/100)
FROM public.sales_person c
LEFT JOIN public.order o ON o.sale_id = c.id
GROUP BY c.id
```

+ b 

```sql
SELECT c.name
FROM public.sales_person c
INNER JOIN public.order o ON o.sale_id = c.id
WHERE o.customer_id <> 2
```

### 2

doạn chương trình sai scope của properties `age`, nếu giữ nguyên scope private thì lớp kế thừa Boss không dùng age, 
trong trường hợp muốn dùng có thể dùng `getter` hoặc đổi scope age thành `protected` hoặc `public`, prefer dùng getter 

```java
class Person {
    protected String name;

    @Getter
    private int age;

    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    @Override
    public String toString() {
        return String.format("%s%s", this.name, this.age);
    }
}

class Boss extends Person {
    public Boss(String name, int age) {
        super(name, age);
    }

    @Override
    public String toString() {
        return String.format("%s%s", this.name, getAge());
    }
}

```

#### 4b

self coding
```go
type S struct {
	X int
	Y string
	// multil property
}

func compare(a, b S) bool {
	if &a == &b {
		return true
	}
	if a.X != b.X {
		return false
	}
	if a.Y != b.Y {
		return false
	}
  // more checking
	return true
}
```

Use a builtin function
```go
reflect.DeepEqual(X, Y)
```

#### 5

+ a: Pub/sub là một kiểu kiếu kiến trúc cung cấp một framework, cơ chế để trao đổi message giữa publishers và subcribers trong đó, publishers là người gửi message và subcribers là người nhận message, mỗi message được gửi trên một topic cụ thể.
- b: Observer là một trong Behavior Patterns cung cấp một cơ chế cho phép đối tượng quan sát thông báo bất kì sự kiện nào xảy ra trên đối tượng này đến các đối tượng quan sát

#### 6

+ Chia nhỏ dữ liệu để query nhiều lần
+ Sử dụng index
+ Tầng API cache data

#### 7

redis có thể dùng để cache dữ liệu, có thể đùng làm bộ đếm, dùng làm một hệ thống pub/sub cũng như dùng đế làm message queue

+ a: process redis-server có 1 thread

- b: Redis lưu dữ liệu trên ram tuy nhiên redis vẫn lưu dữ liệu xuống file để đảm bảo dữ liệu ko bị mất, tuy nhiên trong quá trình backup xuống file có thể dẫn đến redis chiếm hết IO trên server, giải pháp build cơ chế backup định kì cho redis, hoặc tách redis server riêng

#### 8

<img src="https://github.com/qunv/test/blob/d14798b962dc3c4425ac8fd72e9858f9fd3746ca/graduation-slide.jpg">

#### 9
 
 + a: GC có tác dụng thu hồi lại bộ nhớ đã được cấp phát cho các object mà không còn được tham chiếu nữa

#### 10

1. Xứ lý sql injection

+ Lọc từ khóa, ký tự đặc biệt do người dùng nhập vào

- Không cộng chuỗi tạo sql

* Phân quyền trong database

2. Xss

+ Validate data 

* filter và loại bỏ keyword
