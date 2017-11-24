<?php

namespace Tests\Client;

use Curl\Curl;
use PetrLatsko\JsonBundle\Client\CurlClient;
use PetrLatsko\JsonBundle\Client\Exception\ClientParseJsonException;
use PetrLatsko\JsonBundle\Client\Exception\ClientTransportException;
use PHPUnit\Framework\TestCase;

class CurlClientTest extends TestCase
{
    /**
     * @var Curl
     */
    protected static $curl;

    /**
     * @var \PHPUnit_Framework_MockObject_MockObject
     */
    protected static $mock;

    public static function setUpBeforeClass()
    {
        self::$curl = new Curl;
    }

    public function setUp()
    {
        if (!self::$mock)
        {
            self::$mock = $this->getMockBuilder(CurlClient::class)
                               ->enableProxyingToOriginalMethods()
                               ->getMock();
        }
    }

    public function testSetTransport()
    {
        self::$mock->expects($this->once())
             ->method('setTransport')
             ->with($this->identicalTo(self::$curl));

        self::$mock->setTransport(self::$curl);
    }

    public function testGetJson($attribute, $expected, $validJsonString)
    {
    }

    public function validJsonString()
    {
        return [
            ['success', false, '{"data":{"message":"error message","code":"error code"},"success":false}'],
            ['success', true,  '{"data":{"prop":[{"name":"Name","scope":{"foo":0.01,"bar":1.00}}]},"success":true}'],
        ];
    }

    public function invalidJsonString()
    {
        return [
            ['{data:{"prop":[{"name":Name]}},}'],
        ];
    }

}
