require "mustache"
class Simple < Mustache
  def name
    "Chris"
  end

  def value
    10_000
  end

  def taxed_value
    value * 0.6
  end

  def in_ca
    true
  end

  def gravatar(one, render)
    return "llll"
  end
end

Simple.template_file ="template.yml"
file = File.open("template-output.yml","w")
file.write(Simple.render)

