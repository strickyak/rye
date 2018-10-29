must '/true/false/' == '/{a_111}/{z_999}/'.format(a_111=True, z_999=False)

must '/123/456/' == '/{}/{}/'.format(123, 456)

must '/123/456/' == '/{0}/{1}/'.format(123, 456)

must '/"true"/"false"/' == '/{a_111%q}/{z_999%q}/'.format(a_111='true', z_999='false')

must '/"123"/"456"/' == '/{%q}/{%q}/'.format('123', '456')

must '/"123"/"456"/' == '/{0%q}/{1%q}/'.format('123', '456')


must '/red/blue/' == '/{one.color}/{two.color}/'.format(one=dict(color='red'), two=dict(color='blue'))
must '/red/blue/' == '/{outer.one.color}/{outer.two.color}/'.format(outer=dict(one=dict(color='red'), two=dict(color='blue')))
must '/e/l/' == '/{outer.one.color[1]}/{outer.two.color[1]}/'.format(outer=dict(one=dict(color='red'), two=dict(color='blue')))
must '/3/7/' == '/{0[1]}/{0[3]}/'.format([2, 3, 5, 7, 11, 13])
must '/3/7/' == '/{0.primes[1][1]}/{0.primes[3][1]}/'.format(dict(primes=[(0, 2), (0, 3), (0, 5), (0, 7), (0, 11), (0, 13)]))
