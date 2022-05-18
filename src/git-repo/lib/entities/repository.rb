# frozen_string_literal: true

module Entities
  Repository = Struct.new(
    :id,
    :name,
    :full_name,
    :html_url,
    :stargazers_count,
    :ssh_url,
    keyword_init: true
  ) do
    def as_alfred_item
      {
        title: full_name,
        subtitle: format("%s?stars=%s", html_url, stargazers_count),
        icon: 1,
        arg: html_url,
        star: stargazers_count,
        text: {
          copy: ssh_url,
          largetype: full_name
        }
      }
    end
  end
end
