WITH users_count AS (
    SELECT COUNT(*) AS total FROM users 
)

INSERT INTO posts ( 
    user_id, 
    title, 
    slug,  
    excerpt,
    image,
    content, 
    category_id, 
    created_at, 
    updated_at
) 
SELECT 
    ceil(random() * (SELECT total FROM users_count)),
    concat('A Sample Post no.', num),
    concat('a-sample-post-',num), 
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla commodo nisl sed odio hendrerit, sit amet dignissim libero fringilla. Fusce vehicula enim eget mauris suscipit, at porta nibh fermentum.',
    'example.JPG',
    E'### 🌟 Lorem Ipsum 👋

*Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla posuere neque id magna pretium rutrum.* 

###### 🌍 Fames Mauris

📚 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla commodo nisl sed odio hendrerit, sit amet dignissim libero fringilla. Fusce vehicula enim eget mauris suscipit, at porta nibh fermentum. Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed euismod turpis eget nisl molestie, id feugiat libero scelerisque. Curabitur convallis augue eu nisi fringilla, id vestibulum velit finibus. Pellentesque sit amet aliquet justo.

 
###### ☕ Nullam aliquam

🌱 Vivamus in libero varius, feugiat libero ut, pharetra lorem. Sed nec lacinia ante. Nam at mauris non libero bibendum lobortis. Nullam aliquam erat a tellus placerat, non ultricies libero ultrices. Integer ullamcorper, sem nec ultricies auctor, nisi mauris sollicitudin quam, eu eleifend nunc ipsum sed enim.   

**Code**
``` js
var foo = function (bar) {
  return bar++;
};
```

___

###### 🎨 Curabitur convallis  

🎶 Ut at leo eu libero posuere eleifend. Cras congue vestibulum magna, id euismod nisl lacinia non. Donec ac dapibus lectus, eu congue enim. Integer non tellus ipsum. Sed malesuada sapien et odio fermentum, et ultricies purus efficitur.

*Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla posuere neque id magna pretium rutrum.*',
    ceil(random() * 8),
    now() - interval '1 hour 14 minute' * (10000 - num),
    now() - interval '1 hour 14 minute' * (10000 - num)
FROM generate_series(1, 10000, 1) as num;

