# Database

### Insert new user

```sql
INSERT INTO Users (user_id, username, full_name, email, password, avatar, join_at) VALUES (
  "3vctipXOR1x6YQ",
  "ahmadhabibi14",
  "Ahmad Rizky Nusantara Habibi",
  "habi@ternaklinux.com",
  "$2a$10$nHTWg9P1NFORdW27i880NOa04eYahq0HfyUOojFV10o792irQZ4b2",
  "/avatars/ahmadhabibi14.png",
  CURRENT_DATE()
);
```