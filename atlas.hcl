variable "url" {
  type    = string
  default = "postgres://postgres:pass@127.0.0.1:5432/polatdb?sslmode=disable"
}

env "local" {
  # NATIVE KAYNAK: Doğrudan Go Modelleri
  src = "ent://ent/schema"
  
  # HEDEF VERİTABANI
  url = var.url
  
  # NATIVE MULTI-SCHEMA: Atlas'a hangi şemaları yönetmesi gerektiğini söylüyoruz.
  # "public" şemasını da ekleyerek Atlas'ın tüm DB bağlamını anlamasını sağlıyoruz.
  schemas = ["public", "core", "finance"]
  
  # Simülasyon alanı (Multi-schema desteği için)
  dev = "docker://postgres/15/dev"

  migration {
    dir = "file://ent/migrate/migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
